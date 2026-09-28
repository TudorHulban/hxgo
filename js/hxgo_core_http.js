document.addEventListener('DOMContentLoaded', (event) => {
    const CONFIG = Object.freeze({
        HTTP_REQUEST_TIMEOUT: 30000,
        MS_FILE_UPLOAD_RESET_TIMEOUT: 1500,
        CACHE_MAX_ENTRIES: 100,
        CACHE_DEFAULT_TTL: 300
    });

    const HX = Object.freeze({
        GET: 'hx-get',
        POST: 'hx-post',
        UPLOAD: 'hx-upload',
        REDIRECT: 'hx-redirect',
        SWAP: 'hx-swap',
        SEND: 'hx-send',
        REQUIRE: 'hx-require',
        MIN: 'hx-min',
        MAX: 'hx-max',
        VDISABLE: 'hx-vdisable',
        PDISABLE: 'hx-pdisable',
        SHOW: 'hx-show',
        HIDE: 'hx-hide',
        SHOW_ONLOAD: 'hx-show-onload',
        ONCHANGE_ENABLE: 'hx-onchange-enable'
    });

    const VERSION = "0.1.3-core-http";
    console.log("version %s", VERSION);

    // --- Request State ---
    let pendingRequests = new Map();
    let requestIdCounter = 0;

    // --- Event System ---
    const fireEvent = (name, detail = {}, options = {}) => {
        const event = new CustomEvent(`hx:${name}`, {
            detail,
            cancelable: options.cancelable !== false,
            bubbles: true
        });
        return document.dispatchEvent(event);
    };

    // --- Auth Token Management ---
    function getAuthToken() {
        const metaToken = document.querySelector('meta[name="auth-token"]');
        if (metaToken && metaToken.content) return metaToken.content;
        const match = document.cookie.match(/(?:^|;\s*)auth_token=([^;]+)/);
        return match ? match[1] : null;
    }

    function getCSRFToken() {
        const metaToken = document.querySelector('meta[name="csrf-token"]');
        if (metaToken && metaToken.content) return metaToken.content;
        const match = document.cookie.match(/(?:^|;\s*)csrf_token=([^;]+)/);
        return match ? match[1] : null;
    }

    // --- Request Management ---
    function retireRequest(id, responseData = null, extra = {}) {
        if (!id || !pendingRequests.has(id)) return;
        const entry = pendingRequests.get(id);
        clearTimeout(entry.timerId);
        pendingRequests.delete(id);

        fireEvent('afterResponse', {
            requestId: id,
            pendingCount: pendingRequests.size,
            response: responseData,
            element: entry.element,
            endpoint: entry.endpoint,
            method: entry.method,
            ...extra
        });
    }

    // --- DOM Swap Helper ---
    function swapElementById(id, outerHTML, requestId = null) {
        const target = document.getElementById(id);
        if (!target) return;

        fireEvent('beforeSwap', { id, target, content: outerHTML, requestId });
        target.outerHTML = outerHTML;
        const fresh = document.getElementById(id);

        if (fresh) {
            fireEvent('afterSwap', { id, target: fresh, requestId });
            fireEvent('reattachListeners', { element: fresh });
        }
    }

    // --- Process HTML Response ---
    function processHtmlResponse(id, html) {
        if (html.startsWith('unknown endpoint:')) {
            retireRequest(id, null, { error: html });
            fireEvent('error', { message: html, requestId: id });
            return;
        }

        const temp = document.createElement('div');
        temp.innerHTML = html;

        // Check for redirect
        const redirectElement = temp.querySelector('[hx-redirect]');
        if (redirectElement) {
            const redirectUrl = redirectElement.getAttribute('hx-redirect');

            const redirectEvent = fireEvent('redirect', {
                url: redirectUrl,
                requestId: id,
                response: html,
                cancelable: true
            });

            retireRequest(id, html, { redirected: true, redirectUrl });

            if (!redirectEvent.defaultPrevented) {
                window.location.href = redirectUrl;
            }
            return;
        }

        const elements = temp.querySelectorAll('[id]');
        if (elements.length === 0) {
            retireRequest(id, html);
            return;
        }

        elements.forEach(el => swapElementById(el.id, el.outerHTML, id));

        retireRequest(id, html);
    }

    // --- Core HTTP Request Sending ---
    const sendHttpAction = async (element, form) => {
        const id = 'hx' + (++requestIdCounter);

        const endpoint = element.getAttribute(HX.GET) || element.getAttribute(HX.POST);
        const isPost = element.hasAttribute(HX.POST);
        const params = new URLSearchParams();

        // Collect form data + hx-send values for BOTH GET and POST
        const formData = new FormData(form);

        const hxSend = element.getAttribute(HX.SEND);
        if (hxSend) {
            hxSend.split(',').map(id => id.trim()).forEach(id => {
                const el = document.querySelector(id);
                if (!el) return;
                const value = ['INPUT', 'TEXTAREA', 'SELECT'].includes(el.tagName)
                    ? el.value
                    : (el.innerText || el.textContent);
                formData.append(el.id || id, value);
            });
        }

        for (const [key, value] of formData.entries()) {
            if (value instanceof File) continue;
            params.append(key, value);
        }

        // CSRF is only relevant for state-changing requests; keeping it out
        // of GET query strings avoids leaking it into logs/history/Referer.
        if (isPost) {
            const csrf = getCSRFToken();
            if (csrf) params.append('_csrf', csrf);
        }

        // Fire before request - plugins can cancel
        const beforeEvent = fireEvent('beforeRequest', {
            element,
            form,
            requestId: id,
            cancelable: true
        });

        if (beforeEvent.defaultPrevented) {
            fireEvent('requestCancelled', {
                element,
                requestId: id
            });

            fireEvent('afterResponse', {
                requestId: id,
                pendingCount: pendingRequests.size,
                response: null,
                element,
                endpoint,
                method: isPost ? 'POST' : 'GET',
                cancelled: true
            });

            return;
        }

        const controller = new AbortController();
        const timerId = setTimeout(() => {
            const entry = pendingRequests.get(id);
            if (!entry) return;

            entry.timedOut = true;
            pendingRequests.delete(id);
            controller.abort();

            fireEvent('timeout', {
                requestId: id,
                element,
                endpoint,
                method: isPost ? 'POST' : 'GET'
            });

            fireEvent('afterResponse', {
                requestId: id,
                pendingCount: pendingRequests.size,
                response: null,
                element,
                endpoint,
                method: isPost ? 'POST' : 'GET',
                timedOut: true
            });
        }, CONFIG.HTTP_REQUEST_TIMEOUT);

        pendingRequests.set(id, {
            timerId,
            element,
            endpoint,
            method: isPost ? 'POST' : 'GET',
            controller,
            timedOut: false
        });

        params.append('_hx_req_id', id);

        fireEvent('requestSent', {
            requestId: id,
            element,
            endpoint,
            method: isPost ? 'POST' : 'GET'
        });

        fireEvent('afterRequest', {
            requestId: id,
            element,
            endpoint,
            method: isPost ? 'POST' : 'GET'
        });

        try {
            const url = isPost ? endpoint : `${endpoint}?${params.toString()}`;

            const headers = isPost
                ? { 'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8' }
                : {};

            const authToken = getAuthToken();
            if (authToken) headers['Authorization'] = `Bearer ${authToken}`;

            const response = await fetch(url, {
                method: isPost ? 'POST' : 'GET',
                headers,
                body: isPost ? params.toString() : undefined,
                signal: controller.signal
            });

            clearTimeout(timerId);

            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(`HTTP ${response.status} ${errorText}`);
            }

            const html = await response.text();
            processHtmlResponse(id, html);
        } catch (err) {
            // If timeout already handled the request, just return
            if (!pendingRequests.has(id)) return;

            clearTimeout(timerId);
            pendingRequests.delete(id);

            if (err.name === 'AbortError') {
                fireEvent('error', { error: err, element, requestId: id });
                fireEvent('afterResponse', {
                    requestId: id,
                    pendingCount: pendingRequests.size,
                    response: null,
                    element,
                    endpoint,
                    method: isPost ? 'POST' : 'GET',
                    error: err
                });
            } else {
                fireEvent('error', { message: err.message, error: err, element, requestId: id });
                fireEvent('afterResponse', {
                    requestId: id,
                    pendingCount: pendingRequests.size,
                    response: null,
                    element,
                    endpoint,
                    method: isPost ? 'POST' : 'GET',
                    error: err
                });
            }
        }
    };

    // --- Core Upload (HTTP) ---
    const handleUpload = (element, form, file = null) => {
        const beforeEvent = fireEvent('beforeUpload', {
            element,
            form,
            file,
            cancelable: true
        });

        if (beforeEvent.defaultPrevented) {
            fireEvent('uploadCancelled', { element });
            return;
        }

        const endpoint = element.getAttribute(HX.UPLOAD);
        const targetSelectors = element.getAttribute(HX.SWAP);
        const targetElements = targetSelectors
            ? targetSelectors.split(',').map(selector => document.querySelector(selector.trim()))
            : [];
        const redirectUrl = element.getAttribute(HX.REDIRECT);

        if (!file) return;

        const formData = new FormData(form);
        formData.append('file', file);

        const csrfToken = getCSRFToken();
        if (csrfToken) formData.append('_csrf', csrfToken);

        const xhr = new XMLHttpRequest();
        xhr.open('POST', endpoint, true);

        const authToken = getAuthToken();
        if (authToken) xhr.setRequestHeader('Authorization', `Bearer ${authToken}`);

        xhr.upload.addEventListener('progress', (event) => {
            if (event.lengthComputable) {
                const percentComplete = (event.loaded / event.total) * 100;
                fireEvent('uploadProgress', {
                    element,
                    progress: percentComplete
                });
            }
        });

        xhr.onload = () => {
            fireEvent('uploadComplete', { element, status: xhr.status });

            if (xhr.status >= 200 && xhr.status < 300) {
                if (redirectUrl) {
                    const redirectEvent = fireEvent('redirect', {
                        url: redirectUrl,
                        element,
                        cancelable: true
                    });
                    if (!redirectEvent.defaultPrevented) {
                        window.location.href = redirectUrl;
                    }
                } else {
                    const data = xhr.responseText;
                    if (data) {
                        const extractedHTML = parseString(data, targetElements);
                        targetElements.forEach(targetElement => {
                            if (!targetElement) return;
                            const responseElement = extractedHTML.get(targetElement.id);
                            if (responseElement) {
                                swapElementById(targetElement.id, responseElement);
                            }
                        });
                    }
                }
            } else {
                fireEvent('error', {
                    message: xhr.responseText,
                    element,
                    status: xhr.status
                });
            }
        };

        xhr.onerror = () => {
            fireEvent('error', {
                error: xhr.statusText,
                element
            });
        };

        fireEvent('uploadStarted', { element, file });
        xhr.send(formData);
    };

    // --- Event Handlers (Core) ---
    const handleButtonClick = (event) => {
        event.preventDefault();
        const element = event.currentTarget;
        const form = element.closest('form') || document.createElement('form');

        if (element.hasAttribute(HX.UPLOAD)) {
            handleUploadClick(event);
        } else {
            sendHttpAction(element, form);
        }
    };

    const handleSelectChange = (event) => {
        const element = event.currentTarget;
        const form = document.createElement('form');
        const hiddenInput = document.createElement('input');
        hiddenInput.type = 'hidden';
        hiddenInput.name = 'selectedOption';
        hiddenInput.value = element.value;
        form.appendChild(hiddenInput);
        sendHttpAction(element, form);
    };

    const handleUploadClick = (event) => {
        event.preventDefault();
        const element = event.currentTarget;
        const form = element.closest('form') || document.createElement('form');

        let fileInput = form.querySelector('input[type="file"]');
        if (!fileInput) {
            fileInput = document.createElement('input');
            fileInput.type = 'file';
            fileInput.name = 'file';
            fileInput.style.display = 'none';
            form.appendChild(fileInput);
        }

        fileInput.value = '';
        fileInput.addEventListener('change', function onChange() {
            const fileUploaded = fileInput.files[0];
            handleUpload(element, form, fileUploaded);
            setTimeout(() => {
                fileInput.value = '';
            }, CONFIG.MS_FILE_UPLOAD_RESET_TIMEOUT);
        }, { once: true });

        fileInput.click();
    };

    // --- Utility Functions ---
    function parseString(inputString, elements) {
        let parts = inputString.split('|');
        let matchedElementsMap = new Map();

        for (let part of parts) {
            for (let element of elements) {
                if (element && element.id && part.includes(`id="${element.id}"`)) {
                    matchedElementsMap.set(element.id, part);
                    break;
                }
            }
        }
        return matchedElementsMap;
    }

    // --- Event Delegation for Dynamic Elements ---
    document.addEventListener('click', (event) => {
        const element = event.target.closest(
            `button[${HX.GET}], button[${HX.POST}], button[${HX.UPLOAD}], a[${HX.GET}], a[${HX.POST}], a[${HX.UPLOAD}]`
        );
        if (element) {
            event.preventDefault();
            handleButtonClick({ currentTarget: element, preventDefault: () => { } });
        }
    });

    document.addEventListener('change', (event) => {
        const element = event.target.closest(`select[${HX.GET}], select[${HX.POST}]`);
        if (element) {
            handleSelectChange({ currentTarget: element });
        }
    });

    // --- Lifecycle Events ---
    window.addEventListener('beforeunload', () => {
        fireEvent('beforeUnload');
        pendingRequests.forEach(entry => clearTimeout(entry.timerId));
        pendingRequests.clear();
    });

    window.addEventListener('online', () => fireEvent('online'));
    window.addEventListener('offline', () => fireEvent('offline'));

    // --- Expose API for plugins ---
    window.hx = {
        version: VERSION,
        fireEvent,
        sendAction: sendHttpAction,
        sendHttpAction,
        handleUpload,
        swapElementById,
        getAuthToken,
        getCSRFToken,
        config: CONFIG,
        constants: HX
    };

    // Initialize
    fireEvent('init');
});