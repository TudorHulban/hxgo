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

    const VERSION = "0.1.3-core-http-perf";
    console.log("version %s", VERSION);

    // --- Request State ---
    let pendingRequests = new Map();
    let requestIdCounter = 0;

    // --- Debug / Performance ---
    /*
     * HOW TO READ THE [hx-perf] CONSOLE LINES (HTTP VERSION)
     *
     * All values are microseconds (µs): 1,000µs = 1ms. They come from
     * performance.now(). Turn everything off with hx.setDebug(false) or by
     * setting `debug = false` below. Per-element swap detail uses console.debug,
     * so it is only visible when the console level "Verbose" is enabled.
     * Every line starts with a wall-clock stamp [HH:MM:SS.mmm] (browser local time)
     * so it can be lined up with server logs. On one machine the clocks agree.
     *
     * One user action produces these lines, matched by request id (hx1, hx2...):
     *
     *   hx1 -> POST /counter/increment | collect | beforeRequest listeners | prep | fetch() call | client total | size
     *   hx1 <- swap        | rtt (ttfb + body) [= server + net] | parse | swap (N el) | handler total | size
     *   hx1 net detail     | stalled | dns | connect (tls) | waiting | download | protocol | transfer
     *   hx1 paint ≈        | frame-wait + render = X after handler | felt Y (Z without frame-wait)
     *
     * Timeline of one request (left to right is time):
     *
     *   action ─┬─ collect                   FormData, hx-send lookups, URLSearchParams, CSRF
     *           ├─ beforeRequest listeners   plugin handlers of hx:beforeRequest. In the HTTP
     *           │                            version they run AFTER collect (the WS version
     *           │                            runs them first).
     *           ├─ prep                      AbortController, timeout timer, bookkeeping,
     *           │                            requestSent / afterRequest listeners, headers
     *           ├─ fetch() call              ONLY starts the request and returns a promise.
     *           │                            No network, no server.
     *           │   └─ "client total" = the four items above
     *           │
     *           ├──────────── rtt ────────── Just before fetch() -> response BODY fully read.
     *           │   │                        Parse, swap and paint are NOT in it: the value
     *           │   │                        is fixed before any of that runs. If the main
     *           │   │                        thread was busy (e.g. still rendering a previous
     *           │   │                        swap), that wait is inside rtt too.
     *           │   ├─ ttfb                  fetch() -> response headers arrived. Contains
     *           │   │                        connection setup (first request only), sending
     *           │   │                        the request, the server handler, first byte back.
     *           │   └─ body                  headers -> body fully read. The HTML download,
     *           │                            close to 0 for small responses.
     *           │
     *   response processed ─┬─ parse         innerHTML of the response only.
     *                       ├─ swap (N el)   outerHTML replace per [id] element, including
     *                       │                hx:beforeSwap / afterSwap / reattachListeners
     *                       │                listeners. N counts every [id] in the response,
     *                       │                nested ones too (they get swapped twice).
     *                       ├─ handler total all of the response processing, including
     *                       │                hx:afterResponse listeners. Slightly more than
     *                       │                parse + swap.
     *                       └─ paint ≈       the browser's cost after the handler, in two parts:
     *                            frame-wait    handler end -> next rendering opportunity
     *                                          (requestAnimationFrame fires). Mostly display
     *                                          cadence: 0 to 17ms at 60Hz, 0 to 33ms at 30Hz.
     *                                          Nothing to optimise.
     *                            render        rAF -> a task queued from it. Style, layout,
     *                                          paint and commit of the swapped DOM. This is
     *                                          the part that shrinks if the swap gets smaller.
     *
     * "felt" = from the start of sendHttpAction until just after the paint above. It does
     * not include the time the browser spent before calling our input handler. The
     * bracket "(X without frame-wait)" removes frame-wait, which is display cadence and
     * not under your control. Compare that bracketed number before/after an optimisation:
     * it only moves when the app itself got faster or slower.
     *
     * SPLITTING rtt INTO SERVER AND NET
     *   The browser cannot see what the server did, so the server has to say. For HTTP
     *   that is one standard response header, the `app` metric of Server-Timing, with the
     *   duration in MILLISECONDS (a decimal number):
     *
     *     Server-Timing: app;dur=0.187
     *
     *   In Go, set it in the handler just before the body is written (headers go out with
     *   the first write):
     *
     *     w.Header().Set("Server-Timing", fmt.Sprintf("app;dur=%.3f", float64(time.Since(start).Microseconds())/1000))
     *
     *   Then the line shows   rtt 2,700µs (...) = server 187µs + net 2,513µs
     *   with net = rtt - server (clamped at 0, clock coarsening can make it slightly
     *   negative). net holds: connection setup, request and response transfer, the
     *   browser's own queueing before the request left, and main-thread wait.
     *   Start the server's clock as early as possible (middleware, not the handler body)
     *   or request parsing and routing end up in "net".
     *
     * NET DETAIL LINE (needs no server change)
     *   The browser records its own network timeline for every fetch (Resource Timing).
     *   It arrives a moment after the response, so this line prints later than "<-":
     *
     *   - stalled    queued in the browser before DNS / connect / sending (connection pool,
     *                cache check, proxy)
     *   - dns        name lookup. 0 for localhost and for reused connections.
     *   - connect    TCP connect, with the TLS part in brackets. 0 when a keep-alive or
     *                HTTP/2 connection is reused, which is the normal case after the first
     *                request. A TCP handshake involves no server work, so a non-zero value
     *                is about ONE network round trip.
     *   - waiting    request sent -> first byte back. One network round trip plus the
     *                server's time. DevTools calls it "Waiting for server response".
     *   - download   first byte -> last byte.
     *   - protocol   http/1.1, h2, h3. transfer is the bytes on the wire (0 if cached).
     *   - server-timing   every Server-Timing metric the server sent, if any.
     *
     *   rtt minus the sum of these is JavaScript and event-loop overhead inside the page.
     *   Only same-origin requests (or ones with Timing-Allow-Origin) get full detail. A
     *   line is matched to a request by URL and start time, so two identical requests in
     *   flight at the same moment can swap lines.
     *
     * Other lines:
     *   - "<id> cancelled / timed out / failed"   with the time since fetch() started.
     *   - rtt "= server + net" missing            the response had no Server-Timing app
     *                                             metric. One hint is logged.
     *
     * CAVEATS
     *   - performance.now() is coarsened by browsers (Chrome ~100µs, 5µs if the page is
     *     cross-origin isolated, Firefox 1ms by default). A "0µs" means "below the clock
     *     resolution", not zero. Do not trust differences smaller than that.
     *   - The first request after page load is slow (connection setup, cold JIT, cold
     *     server caches). Judge steady state: repeat the action 10 to 20 times and ignore
     *     hx1. If the connection is closed after an idle period, one slow request follows.
     *   - Numbers on 127.0.0.1 exclude real network latency. Add the real RTT of the
     *     link when estimating what remote users see.
     *   - "render" ends at a setTimeout queued from the rAF callback. That is right after
     *     the frame was produced, not a guarantee that pixels reached the screen.
     *   - Open DevTools slows rendering and can disable the browser cache, so paint and
     *     network numbers are pessimistic while it is open.
     */
    let debug = true;
    let warnedNoServerTime = false; // one-time hint when the response has no Server-Timing app metric

    const us = (from, to) => Math.round((to - from) * 1000);
    const fmtUs = (n) => n.toLocaleString('en-US') + 'µs';

    // [HH:MM:SS.mmm] in browser local time, for lining logs up with the server.
    const getTimestamp = () => {
        const now = new Date();
        return now.toTimeString().split(' ')[0] + '.' + String(now.getMilliseconds()).padStart(3, '0');
    };

    const dbg = (...parts) => { if (debug) console.log(`[hx ${getTimestamp()}]`, ...parts); };
    const perfLog = (...parts) => { if (debug) console.log(`[hx-perf ${getTimestamp()}]`, ...parts); };

    // Timing record of one request, filled in as it moves through sendHttpAction
    // and processHtmlResponse. Every t* field is a performance.now() value (ms).
    const newPerf = (tStart) => ({
        tStart,
        tCollected: tStart,
        tBefore: tStart,
        tSend: tStart,        // just before fetch()
        tHeaders: tStart,     // response headers arrived
        tRecv: tStart,        // response body fully read
        tParseStart: tStart,
        tParseEnd: tStart,
        tEnd: tStart,         // end of response processing
        serverUs: null,       // from Server-Timing, null when absent
        swapUs: 0,
        swapped: 0,
        outcome: 'swap'
    });

    // Returns the `app` metric of a Server-Timing header in µs, or null.
    // Header example: "app;dur=0.187, db;dur=0.05" -> 187
    const parseServerTimingUs = (header) => {
        if (!header) return null;

        for (const metric of header.split(',')) {
            const [name, ...params] = metric.split(';').map(part => part.trim());
            if (name !== 'app') continue;

            const dur = params.find(part => part.startsWith('dur='));
            const ms = dur ? Number(dur.slice(4)) : NaN;
            if (Number.isFinite(ms)) return Math.round(ms * 1000);
        }
        return null;
    };

    // One summary line per response, plus a late "paint" line.
    const logResponse = (id, perf, bytes) => {
        if (!debug) return;

        const rttUs = us(perf.tSend, perf.tRecv);
        let rtt = `${fmtUs(rttUs)}`
            + ` (ttfb ${fmtUs(us(perf.tSend, perf.tHeaders))}`
            + ` + body ${fmtUs(us(perf.tHeaders, perf.tRecv))})`;

        if (perf.serverUs !== null) {
            // net = everything that is not the server's own work (see doc comment)
            rtt += ` = server ${fmtUs(perf.serverUs)} + net ${fmtUs(Math.max(0, rttUs - perf.serverUs))}`;
        } else if (!warnedNoServerTime) {
            warnedNoServerTime = true;
            dbg('rtt cannot be split: no "Server-Timing: app;dur=<ms>" response header (the net detail line still shows the browser\'s own breakdown)');
        }

        perfLog(
            `${id} ← ${perf.outcome}`,
            `| rtt ${rtt}`,
            `| parse ${fmtUs(us(perf.tParseStart, perf.tParseEnd))}`,
            `| swap ${fmtUs(perf.swapUs)} (${perf.swapped} el)`,
            `| handler total ${fmtUs(us(perf.tRecv, perf.tEnd))}`,
            `| ${bytes} chars`
        );

        if (perf.swapped > 0) {
            const tDone = perf.tEnd;

            // rAF fires at the next rendering opportunity (frame-wait ends here).
            // The timeout queued from it runs once that frame was produced, so the
            // span in between is style + layout + paint + commit (render).
            requestAnimationFrame(() => {
                const tFrame = performance.now();

                setTimeout(() => {
                    const tPainted = performance.now();
                    const feltUs = us(perf.tStart, tPainted);

                    perfLog(
                        `${id} paint ≈ frame-wait ${fmtUs(us(tDone, tFrame))}`,
                        `+ render ${fmtUs(us(tFrame, tPainted))}`,
                        `= ${fmtUs(us(tDone, tPainted))} after handler`,
                        `| felt ${fmtUs(feltUs)} since action start`,
                        `(${fmtUs(feltUs - us(tDone, tFrame))} without frame-wait)`
                    );
                }, 0);
            });
        }
    };

    // --- Net detail (Resource Timing) ---
    // Requests waiting for their browser-side timing entry: id -> { url, tFetch0 }
    const awaitingNetDetail = new Map();

    // Duration between two Resource Timing marks in µs; 0 when a mark is missing.
    const spanUs = (from, to) => (from > 0 && to >= from) ? Math.round((to - from) * 1000) : 0;

    const logNetDetail = (rt) => {
        if (!debug || rt.initiatorType !== 'fetch') return;

        // Pick the waiting request with the same URL whose start is closest.
        let bestId = null;
        let bestDiff = Infinity;
        for (const [id, waiting] of awaitingNetDetail) {
            if (waiting.url !== rt.name) continue;

            const diff = Math.abs(rt.startTime - waiting.tFetch0);
            if (diff < bestDiff) {
                bestDiff = diff;
                bestId = id;
            }
        }
        if (bestId === null || bestDiff > 200) return;
        awaitingNetDetail.delete(bestId);

        const dnsUs = spanUs(rt.domainLookupStart, rt.domainLookupEnd);
        const connectUs = spanUs(rt.connectStart, rt.connectEnd);
        const tlsUs = spanUs(rt.secureConnectionStart, rt.connectEnd);
        const stalledUs = Math.max(0, us(rt.fetchStart, rt.requestStart) - dnsUs - connectUs);
        const serverTiming = (rt.serverTiming || [])
            .map(metric => `${metric.name} ${fmtUs(Math.round(metric.duration * 1000))}`)
            .join(', ');

        perfLog(
            `${bestId} net detail`,
            `| stalled ${fmtUs(stalledUs)}`,
            `| dns ${fmtUs(dnsUs)}`,
            `| connect ${fmtUs(connectUs)} (tls ${fmtUs(tlsUs)})`,
            `| waiting ${fmtUs(spanUs(rt.requestStart, rt.responseStart))}`,
            `| download ${fmtUs(spanUs(rt.responseStart, rt.responseEnd))}`,
            `| ${rt.nextHopProtocol || 'protocol n/a'}, ${rt.transferSize}B`,
            ...(serverTiming ? [`| server-timing ${serverTiming}`] : [])
        );
    };

    if (typeof PerformanceObserver === 'function') {
        try {
            new PerformanceObserver((list) => list.getEntries().forEach(logNetDetail))
                .observe({ type: 'resource' });
        } catch (_) { /* Resource Timing not supported: no net detail lines */ }
    }

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
        if (!target) {
            dbg(`swap skipped: no element with id "${id}" in the DOM`);
            return;
        }

        const t0 = performance.now();
        fireEvent('beforeSwap', { id, target, content: outerHTML, requestId });
        const t1 = performance.now();
        target.outerHTML = outerHTML;
        const fresh = document.getElementById(id);
        const t2 = performance.now();

        if (fresh) {
            fireEvent('afterSwap', { id, target: fresh, requestId });
            fireEvent('reattachListeners', { element: fresh });
        }

        if (debug) {
            const t3 = performance.now();
            console.debug(
                `[hx-perf ${getTimestamp()}]`,
                `  swap #${id}`,
                `| beforeSwap ${fmtUs(us(t0, t1))}`,
                `| outerHTML ${fmtUs(us(t1, t2))}`,
                `| afterSwap+reattach ${fmtUs(us(t2, t3))}`,
                `| total ${fmtUs(us(t0, t3))}`
            );
        }
    }

    // --- Process HTML Response ---
    // `perf` is the request's timing record (see newPerf); the summary line is
    // logged from the finally block so every exit path is covered.
    function processHtmlResponse(id, html, perf = newPerf(performance.now())) {
        const bytes = html.length;

        try {
            if (html.startsWith('unknown endpoint:')) {
                perf.outcome = 'unknown-endpoint';
                retireRequest(id, null, { error: html });
                fireEvent('error', { message: html, requestId: id });
                return;
            }

            perf.tParseStart = performance.now();
            const temp = document.createElement('div');
            temp.innerHTML = html;
            perf.tParseEnd = performance.now();

            // Check for redirect
            const redirectElement = temp.querySelector('[hx-redirect]');
            if (redirectElement) {
                perf.outcome = 'redirect';
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
                perf.outcome = 'no-swap';
                retireRequest(id, html);
                return;
            }

            const tSwap0 = performance.now();
            elements.forEach(el => swapElementById(el.id, el.outerHTML, id));
            perf.swapUs = us(tSwap0, performance.now());
            perf.swapped = elements.length;

            retireRequest(id, html);
        } finally {
            perf.tEnd = performance.now();
            logResponse(id, perf, bytes);
        }
    }

    // --- Core HTTP Request Sending ---
    const sendHttpAction = async (element, form) => {
        const id = 'hx' + (++requestIdCounter);
        const perf = newPerf(performance.now());

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
        perf.tCollected = performance.now();

        // Fire before request - plugins can cancel
        const beforeEvent = fireEvent('beforeRequest', {
            element,
            form,
            requestId: id,
            cancelable: true
        });

        perf.tBefore = performance.now();

        if (beforeEvent.defaultPrevented) {
            dbg(`${id} cancelled by a beforeRequest listener`);
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
            dbg(`${id} timed out after ${fmtUs(us(perf.tSend, performance.now()))}`);

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

            perf.tSend = performance.now();
            const fetchPromise = fetch(url, {
                method: isPost ? 'POST' : 'GET',
                headers,
                body: isPost ? params.toString() : undefined,
                signal: controller.signal
            });
            const tFetchCalled = performance.now();

            if (debug) {
                perfLog(
                    `${id} → ${isPost ? 'POST' : 'GET'} ${endpoint}`,
                    `| collect ${fmtUs(us(perf.tStart, perf.tCollected))}`,
                    `| beforeRequest listeners ${fmtUs(us(perf.tCollected, perf.tBefore))}`,
                    `| prep ${fmtUs(us(perf.tBefore, perf.tSend))}`,
                    `| fetch() call ${fmtUs(us(perf.tSend, tFetchCalled))}`,
                    `| client total ${fmtUs(us(perf.tStart, tFetchCalled))}`,
                    `| ${isPost ? params.toString().length : url.length} chars`
                );

                // Register for the browser-side timing entry (see logNetDetail).
                for (const [waitingId, waiting] of awaitingNetDetail) {
                    if (perf.tSend - waiting.tFetch0 > 10000) awaitingNetDetail.delete(waitingId); // never matched
                }
                awaitingNetDetail.set(id, {
                    url: new URL(url, window.location.href).href,
                    tFetch0: perf.tSend
                });
            }

            const response = await fetchPromise;
            perf.tHeaders = performance.now();

            clearTimeout(timerId);

            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(`HTTP ${response.status} ${errorText}`);
            }

            const html = await response.text();
            perf.tRecv = performance.now();
            perf.serverUs = parseServerTimingUs(response.headers.get('Server-Timing'));

            processHtmlResponse(id, html, perf);
        } catch (err) {
            // If timeout already handled the request, just return
            if (!pendingRequests.has(id)) return;

            clearTimeout(timerId);
            pendingRequests.delete(id);

            dbg(`${id} failed after ${fmtUs(us(perf.tSend, performance.now()))}: ${err.name}: ${err.message}`);

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
        setDebug: (on) => { debug = Boolean(on); },
        config: CONFIG,
        constants: HX
    };

    // Initialize
    fireEvent('init');
});