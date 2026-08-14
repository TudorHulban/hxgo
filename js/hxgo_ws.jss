document.addEventListener('DOMContentLoaded', (event) => {
    const CONFIG = Object.freeze(
        {
            MS_LOADING_INDICATOR_DELAY: 100,
            MS_DISABLE_TRIGGER_BUTTON: 500,
            MS_DURATION_DISPLAY_POPOVER: 3000,
            MS_DURATION_DISPLAY_ERROR_ALERT: 7000,
            MS_FILE_UPLOAD_RESET_TIMEOUT: 1500,
            WS_REQUEST_TIMEOUT: 30000,
            WS_HEARTBEAT_INTERVAL: 30000,
            WS_RECONNECT_MAX_BACKOFF: 30000,
            WS_RECONNECT_JITTER_MAX: 5000
        }
    );

    const HX = Object.freeze(
        {
            // HTTP Methods
            GET: 'hx-get',
            POST: 'hx-post',
            UPLOAD: 'hx-upload',

            // Element Control
            ENABLE: 'hx-enable',
            DISABLE: 'hx-disable',
            SWAP: 'hx-swap',
            SEND: 'hx-send',
            REDIRECT: 'hx-redirect',

            // Validation
            REQUIRE: 'hx-require',
            MIN: 'hx-min',
            MAX: 'hx-max',
            VDISABLE: 'hx-vdisable',
            PDISABLE: 'hx-pdisable',

            // Visibility/Display
            SHOW: 'hx-show',
            HIDE: 'hx-hide',
            SHOW_ONLOAD: 'hx-show-onload',

            // Event Handlers
            ONCHANGE_ENABLE: 'hx-onchange-enable'
        }
    );

    const VERSION = "0.0.84";

    console.log("version %s", VERSION);

    let loadingIndicatorTimeout;
    let isInitialLoad = true;
    let heartbeatInterval = null;
    let pendingRequests = new Map();
    let requestIdCounter = 0;

    // --- WebSocket State ---
    const WS_URL = (window.location.protocol === 'https:' ? 'wss://' : 'ws://') + window.location.host + '/ws';
    let ws = null;
    let wsBackoff = 1000;
    let wsReconnectTimer = null;
    let wsReconnectAttempts = 0;

    function retireRequest(id) {
        if (!id || !pendingRequests.has(id)) return;
        const entry = pendingRequests.get(id);
        clearTimeout(entry.timerId);
        pendingRequests.delete(id);
        if (pendingRequests.size === 0) {
            hideLoadingIndicator();
        }
    }

    // --- Auth Token Management ---
    function getAuthToken() {
        const metaToken = document.querySelector('meta[name="auth-token"]');
        if (metaToken && metaToken.content) {
            return metaToken.content;
        }
        const match = document.cookie.match(/(?:^|;\s*)auth_token=([^;]+)/);
        return match ? match[1] : null;
    }

    function getCSRFToken() {
        const metaToken = document.querySelector('meta[name="csrf-token"]');
        if (metaToken && metaToken.content) {
            return metaToken.content;
        }
        const match = document.cookie.match(/(?:^|;\s*)csrf_token=([^;]+)/);
        return match ? match[1] : null;
    }

    const loadingIndicator = document.createElement('div');
    loadingIndicator.className = 'loading-indicator';
    loadingIndicator.style.cssText = 'position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:10000;font-family:Material Symbols Outlined;font-size:48px;color:rgba(0,0,0,.8);padding:10px 20px;background-color:hsl(48,85%,26%);border-radius:5px;box-shadow:0 0 10px rgba(0,0,0,.5);display:none;';
    loadingIndicator.textContent = 'sync';
    document.body.appendChild(loadingIndicator);

    function showLoadingIndicator() {
        loadingIndicatorTimeout = setTimeout(() => {
            loadingIndicator.style.display = 'block';
        }, CONFIG.MS_LOADING_INDICATOR_DELAY);
    }

    function hideLoadingIndicator() {
        clearTimeout(loadingIndicatorTimeout);
        loadingIndicator.style.display = 'none';
    }

    const baseElement = document.createElement('base');
    baseElement.href = `${window.location.origin}/`;
    document.head.appendChild(baseElement);

    // --- DOM swap helper ---
    // Single source of truth for "replace element by id, then reattach listeners
    // on exactly that element's subtree". Both the WS message handler and the
    // upload response handler funnel through this so listener (re)binding is
    // driven only by the ids the backend actually told us about — no blanket
    // subtree observation, and no double-binding.
    function swapElementById(id, outerHTML) {
        const target = document.getElementById(id);
        if (!target) return;
        target.outerHTML = outerHTML;
        const fresh = document.getElementById(id);
        if (fresh) reattachEventListeners(fresh);
    }

    function initWs() {
        if (ws?.readyState === WebSocket.CONNECTING || ws?.readyState === WebSocket.OPEN) return;

        const authToken = getAuthToken();
        let wsUrl = WS_URL;
        if (authToken) {
            wsUrl += (wsUrl.includes('?') ? '&' : '?') + 'token=' + encodeURIComponent(authToken);
        }

        try {
            ws = new WebSocket(wsUrl);
        } catch (e) {
            console.error('Failed to create WebSocket:', e);
            scheduleReconnect();
            return;
        }

        ws.onopen = () => {
            console.log('WebSocket connected');
            wsBackoff = 1000;
            wsReconnectAttempts = 0;

            if (heartbeatInterval) {
                clearInterval(heartbeatInterval);
            }
            heartbeatInterval = setInterval(() => {
                if (ws && ws.readyState === WebSocket.OPEN) {
                    ws.send('ping');
                }
            }, CONFIG.WS_HEARTBEAT_INTERVAL);
        };

        ws.onmessage = (event) => {
            let html = event.data;

            if (html === 'pong') {
                return;
            }

            let responseId = null;
            const commentMatch = html.match(/^<!--\s*_hx_req_id:\s*(\S+)\s*-->\s*/);
            if (commentMatch) {
                responseId = commentMatch[1];
                html = html.slice(commentMatch[0].length);
            }

            if (html.startsWith('unknown endpoint:')) {
                retireRequest(responseId);
                showErrorAlert(html);
                return;
            }

            const temp = document.createElement('div');
            temp.innerHTML = html;

            // Check for hx-redirect (same as file upload)
            const redirectElement = temp.querySelector('[hx-redirect]');
            if (redirectElement) {
                const redirectUrl = redirectElement.getAttribute('hx-redirect');
                if (redirectUrl) {
                    window.location.href = redirectUrl;
                    retireRequest(responseId);
                    return;
                }
            }

            const elements = temp.querySelectorAll('[id]');
            if (elements.length === 0) {
                retireRequest(responseId);
                return;
            }

            // Swap strictly by the ids present in this response. This is the
            // only place that binds listeners for WS-driven updates.
            elements.forEach(el => swapElementById(el.id, el.outerHTML));

            retireRequest(responseId);
        };

        ws.onerror = (err) => {
            console.error('WS error:', err);
        };

        ws.onclose = (event) => {
            pendingRequests.forEach(entry => clearTimeout(entry.timerId));
            pendingRequests.clear();
            hideLoadingIndicator();

            console.log('WebSocket closed:', event.code, event.reason);
            if (heartbeatInterval) {
                clearInterval(heartbeatInterval);
                heartbeatInterval = null;
            }
            scheduleReconnect();
        };
    }

    function scheduleReconnect() {
        if (wsReconnectTimer) return;

        const baseDelay = Math.min(wsBackoff, CONFIG.WS_RECONNECT_MAX_BACKOFF);
        const jitter = Math.random() * CONFIG.WS_RECONNECT_JITTER_MAX;
        const delay = baseDelay + jitter;

        console.log(`Scheduling reconnect in ${delay}ms (attempt ${wsReconnectAttempts + 1})`);

        wsReconnectTimer = setTimeout(() => {
            wsReconnectTimer = null;
            wsReconnectAttempts++;
            initWs();
        }, delay);
        wsBackoff = Math.min(wsBackoff * 2, CONFIG.WS_RECONNECT_MAX_BACKOFF);
    }

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

    function showPopover(element, message) {
        const popover = document.createElement('div');
        popover.className = 'popover';
        popover.textContent = message;
        document.body.appendChild(popover);

        const rect = element.getBoundingClientRect();
        popover.style.left = `${rect.left + window.scrollX}px`;
        popover.style.top = `${rect.bottom + window.scrollY}px`;
        popover.style.zIndex = '1000';

        setTimeout(() => {
            if (popover.parentNode) {
                document.body.removeChild(popover);
            }
        }, CONFIG.MS_DURATION_DISPLAY_POPOVER);
    }

    function showErrorAlert(message) {
        const errorAlertBox = document.createElement('div');
        errorAlertBox.className = 'error-alert';
        errorAlertBox.textContent = message;

        document.body.appendChild(errorAlertBox);

        setTimeout(() => {
            if (errorAlertBox.parentNode) {
                document.body.removeChild(errorAlertBox);
            }
        }, CONFIG.MS_DURATION_DISPLAY_ERROR_ALERT);
    }

    function toggleElements(attribute, shouldEnable, triggeringElement) {
        if (!attribute) return;

        const ids = attribute.split(',').map(id => id.trim());
        ids.forEach(
            id => {
                const element = document.querySelector(id);
                if (element) {
                    if (attribute.startsWith(HX.ONCHANGE_ENABLE)) {
                        element.disabled = triggeringElement.disabled;
                    } else {
                        element.disabled = !shouldEnable;
                    }
                }
            },
        );
    }

    const validateLength = (element, showPopup = true) => {
        const minLength = element.getAttribute(HX.MIN) ? parseInt(element.getAttribute(HX.MIN), 10) : null;
        const maxLength = element.getAttribute(HX.MAX) ? parseInt(element.getAttribute(HX.MAX), 10) : null;
        const targetDisableId = element.getAttribute(HX.VDISABLE);
        const targetElement = targetDisableId ? document.querySelector(targetDisableId) : null;
        const valueLength = element.value.length;

        let isValid = true;

        if (minLength !== null && valueLength < minLength) {
            if (showPopup && !isInitialLoad) {
                const elementName = element.name ? element.name.charAt(0).toUpperCase() + element.name.slice(1) : "Element";
                showPopover(element, `${elementName}: Minimum length is ${minLength} characters.`);
            }
            isValid = false;
        }
        if (maxLength !== null && valueLength > maxLength) {
            if (showPopup && !isInitialLoad) {
                const elementName = element.name ? element.name.charAt(0).toUpperCase() + element.name.slice(1) : "Element";
                showPopover(element, `${elementName}: Maximum length is ${maxLength} characters.`);
            }
            isValid = false;
        }

        if (targetElement) {
            targetElement.disabled = !isValid;
        }

        return isValid;
    };

    function validatePasswords(elementPw1, elementPw2, targetElement, showMatchPopup = false) {
        const validLengthPw1 = validateLength(elementPw1, false);
        const validLengthPw2 = validateLength(elementPw2, false);
        const passwordsMatch = elementPw1.value === elementPw2.value;
        const bothFilled = elementPw1.value.length > 0 && elementPw2.value.length > 0;

        if (showMatchPopup && bothFilled && !passwordsMatch && !isInitialLoad) {
            showPopover(elementPw2, 'Passwords must match');
            targetElement.disabled = true;
            return false;
        }

        const isValid = validLengthPw1 && validLengthPw2 && passwordsMatch && bothFilled;
        targetElement.disabled = !isValid;
        return isValid;
    }

    const validateRequirements = (element, form = null) => {
        const requireAttr = element.getAttribute(HX.REQUIRE);
        if (requireAttr) {
            const requiredIds = requireAttr.split(',');

            for (let id of requiredIds) {
                const requiredElement = document.querySelector(id.trim());
                if (requiredElement && !requiredElement.value) {
                    if (!isInitialLoad) {
                        showPopover(requiredElement, 'This field is required.');
                    }
                    return false;
                }
            }
        }

        let allValid = true;

        const lengthElements = form.querySelectorAll(`[${HX.MIN}], [${HX.MAX}]`);
        lengthElements.forEach(
            (el) => {
                if (!validateLength(el, true)) {
                    allValid = false;
                }
            },
        );

        const passwordGroups = new Set();

        form.querySelectorAll(`[${HX.PDISABLE}]`).forEach(
            (el) => {
                const pdisableValue = el.getAttribute(HX.PDISABLE);

                if (!passwordGroups.has(pdisableValue)) {
                    passwordGroups.add(pdisableValue);

                    const ids = pdisableValue.split(',').map(id => id.trim());
                    if (ids.length === 3) {
                        const [idPw1, idPw2, idTarget] = ids;
                        const pw1 = document.querySelector(idPw1);
                        const pw2 = document.querySelector(idPw2);
                        const target = document.querySelector(idTarget);

                        if (pw1 && pw2 && target && !validatePasswords(pw1, pw2, target, true)) {
                            allValid = false;
                        }
                    }
                }
            },
        );

        return allValid;
    };

    // --- WebSocket Actions ---

    const sendWsAction = (element, form) => {
        if (!ws || ws.readyState !== WebSocket.OPEN) {
            showErrorAlert('Not connected to server. Reconnecting...');
            scheduleReconnect();
            return;
        }

        const hxEnable = element.getAttribute(HX.ENABLE);
        if (hxEnable) toggleElements(hxEnable, true);

        const hxDisable = element.getAttribute(HX.DISABLE);
        if (hxDisable) toggleElements(hxDisable, false);

        const endpoint = element.getAttribute(HX.GET) || element.getAttribute(HX.POST);
        const isPost = element.hasAttribute(HX.POST);

        if (isPost && !validateRequirements(element, form)) {
            return;
        }

        const params = new URLSearchParams();

        if (isPost) {
            const formData = new FormData(form);

            const hxSend = element.getAttribute(HX.SEND);
            if (hxSend) {
                hxSend.split(',').map(id => id.trim()).forEach(id => {
                    const el = document.querySelector(id);
                    if (!el) return;
                    const value = ['INPUT', 'TEXTAREA', 'SELECT'].includes(el.tagName)
                        ? el.value
                        : (el.innerText || el.textContent);
                    formData.append(el.id, value);
                });
            }

            for (const [key, value] of formData.entries()) {
                if (value instanceof File) continue; // uploads go via handleUpload/XHR
                params.append(key, value);
            }
        }

        const csrf = getCSRFToken();
        if (csrf) params.append('_csrf', csrf);

        element.disabled = true;
        setTimeout(() => { element.disabled = false; }, CONFIG.MS_DISABLE_TRIGGER_BUTTON);

        const id = 'hx' + (++requestIdCounter);
        const timerId = setTimeout(() => {
            if (pendingRequests.has(id)) {
                pendingRequests.delete(id);
                if (pendingRequests.size === 0) {
                    hideLoadingIndicator();
                }
                showErrorAlert('Request timed out');
            }
        }, CONFIG.WS_REQUEST_TIMEOUT);

        pendingRequests.set(id, { timerId, element });
        if (pendingRequests.size === 1) {
            showLoadingIndicator();
        }

        params.append('_hx_req_id', id);

        // verb + endpoint on the first line, encoded body on the second
        const wire = `${isPost ? 'POST' : 'GET'} ${endpoint}\n${params.toString()}`;

        try {
            ws.send(wire);
        } catch (e) {
            element.disabled = false;
            if (pendingRequests.has(id)) {
                clearTimeout(pendingRequests.get(id).timerId);
                pendingRequests.delete(id);
                if (pendingRequests.size === 0) {
                    hideLoadingIndicator();
                }
            }
            showErrorAlert('Failed to send message: ' + e.message);
        }
    };

    // --- Upload (HTTP only) ---

    const handleUpload = (element, form, file = null) => {
        const hxEnable = element.getAttribute(HX.ENABLE);
        if (hxEnable) {
            toggleElements(hxEnable, true);
        }

        const hxDisable = element.getAttribute(HX.DISABLE);
        if (hxDisable) {
            toggleElements(hxDisable, false);
        }

        const endpoint = element.getAttribute(HX.UPLOAD);
        const targetSelectors = element.getAttribute(HX.SWAP);
        const targetElements = targetSelectors ? targetSelectors.split(',').map(selector => document.querySelector(selector.trim())) : [];
        const redirectUrl = element.getAttribute(HX.REDIRECT);

        if (!validateRequirements(element, form)) {
            return;
        }

        if (!file) {
            return;
        }

        const formData = new FormData(form);
        formData.append('file', file);

        const csrfToken = getCSRFToken();
        if (csrfToken) {
            formData.append('_csrf', csrfToken);
        }

        const xhr = new XMLHttpRequest();
        xhr.open('POST', endpoint, true);

        xhr.upload.addEventListener(
            'progress',
            (event) => {
                if (event.lengthComputable) {
                    const percentComplete = (event.loaded / event.total) * 100;
                    console.log(`File upload progress: ${percentComplete}%`);
                    const progressEvent = new CustomEvent('file-upload-progress', { detail: percentComplete });
                    element.dispatchEvent(progressEvent);
                }
            }
        );

        xhr.onload = () => {
            hideLoadingIndicator();
            if (xhr.status >= 200 && xhr.status < 300) {
                if (redirectUrl) {
                    window.location.href = redirectUrl;
                } else {
                    const data = xhr.responseText;
                    if (data) {
                        let extractedHTML = parseString(data, targetElements);
                        // Swap strictly by the ids resolved above — same helper
                        // the WS handler uses, so listener (re)binding stays
                        // driven only by known ids, with no double-binding.
                        targetElements.forEach(
                            (targetElement) => {
                                if (!targetElement) return;
                                const responseElement = extractedHTML.get(targetElement.id);
                                if (responseElement) {
                                    swapElementById(targetElement.id, responseElement);
                                }
                            },
                        );
                    }
                }
            } else {
                showErrorAlert(xhr.responseText);
            }
        };

        xhr.onerror = () => {
            hideLoadingIndicator();
            console.error('Error:', xhr.statusText);
            showErrorAlert('Upload failed: ' + xhr.statusText);
        };

        element.disabled = true;
        setTimeout(() => {
            element.disabled = false;
        }, CONFIG.MS_DISABLE_TRIGGER_BUTTON);

        showLoadingIndicator();
        xhr.send(formData);
    };

    // --- Event Handlers ---

    const handleButtonClick = (event) => {
        event.preventDefault();
        const element = event.currentTarget;
        const form = element.closest('form') || document.createElement('form');

        if (element.hasAttribute(HX.UPLOAD)) {
            handleUploadClick(event);
        } else {
            sendWsAction(element, form);
        }

        handleHxShowHide(element, HX.SHOW);
        handleHxShowHide(element, HX.HIDE);
    };

    const handleSelectChange = (event) => {
        const element = event.currentTarget;
        const form = document.createElement('form');

        const hiddenInput = document.createElement('input');
        hiddenInput.type = 'hidden';
        hiddenInput.name = 'selectedOption';
        hiddenInput.value = element.value;
        form.appendChild(hiddenInput);

        sendWsAction(element, form);
    };

    const handleUploadClick = (event) => {
        event.preventDefault();
        const element = event.currentTarget;
        const form = element.closest('form') || document.createElement('form');

        if (!validateRequirements(element, form)) {
            return;
        }

        let fileInput = form.querySelector('input[type="file"]');
        if (!fileInput) {
            fileInput = document.createElement('input');
            fileInput.type = 'file';
            fileInput.name = 'file';
            fileInput.style.display = 'none';
            form.appendChild(fileInput);
        }

        // Reset the value first so selecting the same file again still fires
        // 'change', then always reopen the dialog on click (fixes stale
        // re-submission of the previously chosen file on repeat clicks).
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

    const handleDblClickClear = (event) => {
        if (event.target.tagName.toLowerCase() === 'button' || event.target.getElementsByTagName('button').length > 0) {
            return;
        }

        const searchItem = event.currentTarget;
        const elInputs = searchItem.querySelectorAll('input, select');

        for (let input of elInputs) {
            if (input.tagName.toLowerCase() === 'input') {
                input.value = '';
            } else if (input.tagName.toLowerCase() === 'select') {
                input.selectedIndex = 0;
            }
        }
    };

    function applyShowHide(targetSelector, timeToShow, cssTransitionClass, isShow) {
        const targetElement = document.querySelector(targetSelector);
        if (!targetElement) return;

        if (cssTransitionClass) {
            targetElement.classList.add(cssTransitionClass);
        }

        if (isShow) {
            targetElement.style.display = 'block';
            if (timeToShow > 0) {
                setTimeout(() => {
                    targetElement.style.display = 'none';
                    if (cssTransitionClass) {
                        targetElement.classList.remove(cssTransitionClass);
                    }
                }, timeToShow);
            }
        } else {
            targetElement.style.display = 'none';
            if (cssTransitionClass) {
                targetElement.classList.remove(cssTransitionClass);
            }
        }
    }

    function handleHxShowHide(element, attributeName) {
        const attributeValue = element.getAttribute(attributeName);
        if (!attributeValue) return;

        const targets = attributeValue.split(',').map(target => target.trim());
        targets.forEach(
            targetSpec => {
                const parts = targetSpec.split(':');
                if (parts.length >= 1) {
                    const targetSelector = parts[0];
                    const timeToShow = parts.length > 1 ? parseInt(parts[1], 10) : 0;
                    const cssTransitionClass = parts.length > 2 ? parts[2] : '';
                    applyShowHide(targetSelector, timeToShow, cssTransitionClass, attributeName === HX.SHOW);
                }
            },
        );
    }

    function handleHxShowOnLoad(element, attributeName) {
        const attributeValue = element.getAttribute(attributeName);
        if (!attributeValue) return;

        const targets = attributeValue.split(',').map(target => target.trim());
        targets.forEach(
            targetSpec => {
                const parts = targetSpec.split(':');
                if (parts.length >= 1) {
                    const targetSelector = parts[0];
                    const timeToShow = parts.length > 1 ? parseInt(parts[1], 10) : 0;
                    const cssTransitionClass = parts.length > 2 ? parts[2] : '';
                    applyShowHide(targetSelector, timeToShow, cssTransitionClass, attributeName === HX.SHOW_ONLOAD);
                }
            },
        );
    }

    // Binds handlers on a single freshly-inserted element and its descendants.
    // This is the ONLY reattachment path in the app: it's invoked explicitly
    // by swapElementById() with the exact id the backend told us changed.
    // There is no MutationObserver watching the whole document for this —
    // that would rebind listeners on every DOM change (popovers, alerts,
    // loading indicator) for content we already know about, and it was the
    // root cause of elements getting double-bound.
    const reattachEventListeners = (element) => {
        if (!element) return;

        if (element.matches(`[${HX.SHOW}], [${HX.HIDE}]`)) {
            handleHxShowHide(element, HX.SHOW);
            handleHxShowHide(element, HX.HIDE);
        }

        if (element.matches(`button[${HX.GET}], button[${HX.POST}], button[${HX.UPLOAD}], a[${HX.GET}], a[${HX.POST}], a[${HX.UPLOAD}]`)) {
            element.addEventListener('click', handleButtonClick);
        }

        if (element.matches(`select[${HX.GET}], select[${HX.POST}]`)) {
            element.addEventListener('change', handleSelectChange);
        }

        if (element.id === 'items-search') {
            element.addEventListener('dblclick', handleDblClickClear);
        }

        element.querySelectorAll(`[${HX.SHOW}], [${HX.HIDE}]`).forEach(
            el => {
                handleHxShowHide(el, HX.SHOW);
                handleHxShowHide(el, HX.HIDE);
            },
        );

        element.querySelectorAll(`button[${HX.GET}], button[${HX.POST}], button[${HX.UPLOAD}], a[${HX.GET}], a[${HX.POST}], a[${HX.UPLOAD}]`).forEach(
            el => {
                el.addEventListener('click', handleButtonClick);
            },
        );

        element.querySelectorAll(`select[${HX.GET}], select[${HX.POST}]`).forEach(
            el => {
                el.addEventListener('change', handleSelectChange);
            },
        );

        const searchItem = element.querySelector('#items-search');
        if (searchItem) {
            searchItem.addEventListener('dblclick', handleDblClickClear);
        }
    };

    document.querySelectorAll(`[${HX.SHOW}], [${HX.HIDE}]`).forEach(
        el => {
            handleHxShowHide(el, HX.SHOW);
            handleHxShowHide(el, HX.HIDE);
        },
    );

    document.querySelectorAll(`[${HX.SHOW_ONLOAD}]`).forEach(
        el => {
            handleHxShowOnLoad(el, HX.SHOW_ONLOAD);
        },
    );

    const processedGroups = new Set();

    document.querySelectorAll(`[${HX.PDISABLE}]`).forEach(
        (element) => {
            const pdisableValue = element.getAttribute(HX.PDISABLE);
            if (processedGroups.has(pdisableValue)) return;

            processedGroups.add(pdisableValue);

            const ids = pdisableValue.split(',').map(id => id.trim());
            if (ids.length !== 3) return;

            const [pw1Id, pw2Id, targetId] = ids;
            const pw1Element = document.querySelector(pw1Id);
            const pw2Element = document.querySelector(pw2Id);
            const targetElement = document.querySelector(targetId);

            if (!pw1Element || !pw2Element || !targetElement) return;

            targetElement.disabled = true;

            const validateBothTyping = () => {
                validatePasswords(pw1Element, pw2Element, targetElement, false);
                isInitialLoad = false;
            };

            const validateBothBlur = () => {
                validatePasswords(pw1Element, pw2Element, targetElement, true);
                isInitialLoad = false;
            };

            pw1Element.addEventListener('input', validateBothTyping);
            pw2Element.addEventListener('input', validateBothTyping);
            pw1Element.addEventListener('blur', validateBothBlur);
            pw2Element.addEventListener('blur', validateBothBlur);

            validatePasswords(pw1Element, pw2Element, targetElement, false);
        },
    );

    const buttonElements = document.querySelectorAll(`button[${HX.GET}], button[${HX.POST}], button[${HX.UPLOAD}], a[${HX.GET}], a[${HX.POST}], a[${HX.UPLOAD}]`);
    buttonElements.forEach(el => el.addEventListener('click', handleButtonClick));

    const elementsSelect = document.querySelectorAll(`select[${HX.GET}], select[${HX.POST}]`);
    elementsSelect.forEach(el => el.addEventListener('change', handleSelectChange));

    const searchItems = document.getElementById('items-search');
    if (searchItems) {
        searchItems.addEventListener('dblclick', handleDblClickClear);
    }

    document.querySelectorAll(`[${HX.MIN}], [${HX.MAX}]`).forEach(
        (element) => {
            element.addEventListener('change', () => {
                validateLength(element);
                isInitialLoad = false;
            });
        },
    );

    document.querySelectorAll(`[${HX.ONCHANGE_ENABLE}]`).forEach(
        element => {
            element.addEventListener('change', function () {
                const hxChangeEnable = this.getAttribute(HX.ONCHANGE_ENABLE);
                toggleElements(`${HX.ONCHANGE_ENABLE},${hxChangeEnable}`, false, this);
                isInitialLoad = false;
            });
        },
    );

    // NOTE: no MutationObserver here. All dynamic DOM swaps in this app
    // (WS message handling, upload response handling) go through
    // swapElementById(), which reattaches listeners on exactly the element
    // that changed. If a future feature injects hx-* markup through some
    // other path (e.g. a third-party script), call reattachEventListeners()
    // on that specific node explicitly rather than reintroducing a
    // document-wide observer.

    function openInNewTab(event) {
        event.preventDefault();
        window.open(event.currentTarget.href, '_blank');
    };

    const links = document.getElementsByClassName('ntab');
    Array.from(links).forEach(
        link => {
            link.addEventListener('click', openInNewTab);
        },
    );

    window.addEventListener('beforeunload', () => {
        if (heartbeatInterval) {
            clearInterval(heartbeatInterval);
            heartbeatInterval = null;
        }

        if (ws) {
            ws.close(1000, 'Page unloading');
            ws = null;
        }

        if (wsReconnectTimer) {
            clearTimeout(wsReconnectTimer);
            wsReconnectTimer = null;
        }

        hideLoadingIndicator();
    });

    document.addEventListener('visibilitychange', () => {
        if (document.visibilityState === 'visible') {
            if (!ws || ws.readyState === WebSocket.CLOSED || ws.readyState === WebSocket.CLOSING) {
                console.log('Tab became visible, reconnecting WebSocket');
                initWs();
            }
        }
    });

    window.addEventListener('online', () => {
        console.log('Network came back online, reconnecting WebSocket');
        if (!ws || ws.readyState !== WebSocket.OPEN) {
            initWs();
        }
    });

    window.addEventListener('offline', () => {
        console.log('Network went offline');
    });

    initWs();
});