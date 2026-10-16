document.addEventListener('DOMContentLoaded', (event) => {
    const CONFIG = Object.freeze({
        WS_REQUEST_TIMEOUT: 30000,
        WS_HEARTBEAT_INTERVAL: 30000,
        WS_RECONNECT_MAX_BACKOFF: 30000,
        WS_RECONNECT_JITTER_MAX: 5000,
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

    const VERSION = "0.1.4-core-ws-perf";
    console.log("version %s", VERSION);

    // --- WebSocket State ---
    const WS_URL = (window.location.protocol === 'https:' ? 'wss://' : 'ws://') + window.location.host + '/ws';
    let ws = null;
    let wsBackoff = 1000;
    let wsReconnectTimer = null;
    let wsReconnectAttempts = 0;
    let heartbeatInterval = null;
    let pendingRequests = new Map();
    let requestIdCounter = 0;

    // --- Debug / Performance ---
    /*
     * HOW TO READ THE [hx-perf] CONSOLE LINES
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
     *   hx1 -> POST /today | beforeRequest listeners | collect+prep | ws.send() | client total | size
     *   hx1 <- swap        | rtt [= server + net | ≈ ping-baseline + server-side] | parse | swap (N el) | handler total | size
     *   hx1 paint ≈        | frame-wait + render = X after handler | felt Y since action start
     *
     * Timeline of one request (left to right is time):
     *
     *   action ─┬─ beforeRequest listeners   plugin handlers of hx:beforeRequest
     *           ├─ collect+prep              FormData, hx-send lookups, URLSearchParams,
     *           │                            CSRF, timeout timer, wire string
     *           ├─ ws.send()                 ONLY queues the frame in the browser buffer.
     *           │                            It does not include the network or the server.
     *           │   └─ "client total" = the three items above
     *           │
     *           ├──────────── rtt ────────── Just before ws.send() -> first line of onmessage.
     *           │                            = server handler + server Send + transport.
     *           │                            With a server-reported time it is split into
     *           │                            "= server X + net Y" (see SPLITTING rtt below).
     *           │                            Parse, swap and paint are NOT in it: the value is
     *           │                            fixed before any of that runs. It is not called
     *           │                            "net-rtt" because it includes the server's work.
     *           │                            If the main thread was busy when the frame
     *           │                            arrived (e.g. still rendering a previous swap),
     *           │                            that wait is inside rtt too.
     *           │
     *   onmessage starts ─┬─ parse           innerHTML of the response only. The header strip
     *                     │                  (regex + slice) is microseconds and is only part
     *                     │                  of "handler total".
     *                     ├─ swap (N el)     outerHTML replace per [id] element, including
     *                     │                  hx:beforeSwap / afterSwap / reattachListeners
     *                     │                  listeners. N counts every [id] in the response,
     *                     │                  nested ones too (they get swapped twice).
     *                     ├─ handler total   all of onmessage, including hx:afterResponse
     *                     │                  listeners. Slightly more than parse + swap.
     *                     └─ paint ≈         the browser's cost after the handler, in two parts:
     *                          frame-wait      handler end -> next rendering opportunity
     *                                          (requestAnimationFrame fires). Mostly vsync
     *                                          wait: 0 to 17ms at 60Hz, nothing to optimise.
     *                          render          rAF -> a task queued from it. Style, layout,
     *                                          paint and commit of the swapped DOM. This is
     *                                          the part that shrinks if the swap gets smaller.
     *
     * "felt" = from the start of sendWsAction until just after the paint above. It does
     * not include the time the browser spent before calling our input handler. The
     * bracket "(X without frame-wait)" removes frame-wait, which is display cadence and
     * not under your control. Compare that bracketed number before/after an optimisation:
     * it only moves when the app itself got faster or slower.
     *
     * SPLITTING rtt INTO SERVER AND NET
     *   The browser cannot see what the server did, so the server has to say. It adds its
     *   own duration, in integer µs, to the response header comment:
     *
     *     <!-- _hx_req_id: hx4 srv=187 -->
     *
     *   Then the line shows   rtt 2,700µs = server 187µs + net 2,513µs
     *   with net = rtt - srv (clamped at 0, clock coarsening can make it slightly negative).
     *   Without srv= the browser can still ESTIMATE the split from a ping baseline. After
     *   every connect (and whenever you run hx.pingTest()) it sends 10 pings and keeps the
     *   median pong rtt as the "ping-baseline". Each response then prints:
     *
     *     rtt 2,400µs ≈ ping-baseline 600µs + server-side 1,800µs
     *
     *   server-side = rtt - baseline = everything the request did that a ping did not:
     *   dispatch, parsing, auth, the handler and the write path. It is an estimate, because
     *   a ping may take a shorter path through the server than a request does. If your
     *   own handler timing is far below "server-side", the gap is in that surrounding
     *   code, not in the handler. An exact srv= always wins over the estimate. Recalibrate
     *   with hx.pingTest() if the baseline looks stale.
     *
     *   Where the server starts and stops its clock decides what "net" contains:
     *     - Start when the frame is READ, not when the handler starts. Otherwise dispatch
     *       and queueing time inside the server hides in "net".
     *     - Stop as late as possible, right before the socket write.
     *   net then holds: transport both ways, the write itself, and any browser
     *   main-thread wait before onmessage ran. On loopback that is about the heartbeat
     *   rtt (hx.pingTest() gives a better baseline). A big "net" on one request with a
     *   normal "server" means the delay was NOT your backend.
     *
     * Other lines:
     *   - "connected ... in X"      WebSocket handshake time, once per connection.
     *   - "pong, heartbeat rtt X"   ping/pong over the same socket. The server does no
     *                               work for it, so it approximates the pure transport
     *                               cost. (request rtt - heartbeat rtt) is roughly the
     *                               time your server handler took.
     *   - hx.pingTest(n, gapMs)     console helper: n pings, one at a time. Prints
     *                               min / median / max transport rtt and stores the
     *                               median as the rtt baseline (it also runs once after
     *                               every connect). Flat pings but a
     *                               jumpy request rtt means the jitter is in the server's
     *                               request path. Jumpy pings mean transport or the
     *                               server's read loop.
     *   - "buffered NB"             ws.bufferedAmount right after send(). A value that
     *                               keeps growing means the connection cannot keep up.
     *   - rtt "n/a"                 response had no request id (server push) or the
     *                               request already timed out.
     *
     * CAVEATS
     *   - performance.now() is coarsened by browsers (Chrome ~100µs, 5µs if the page is
     *     cross-origin isolated, Firefox 1ms by default). A "0µs" means "below the clock
     *     resolution", not zero. Do not trust differences smaller than that.
     *   - The first request after page load is slow (cold JIT, cold server caches).
     *     Judge steady state: repeat the action 10 to 20 times and ignore hx1.
     *   - Numbers on 127.0.0.1 exclude real network latency. Add the real RTT of the
     *     link when estimating what remote users see.
     *   - "render" ends at a setTimeout queued from the rAF callback. That is right after
     *     the frame was produced, not a guarantee that pixels reached the screen.
     *   - Open DevTools slows rendering, so paint numbers are pessimistic while it is open.
     */
    let debug = true;
    let wsConnectStart = 0;
    let pingSentAt = 0;

    const us = (from, to) => Math.round((to - from) * 1000);
    const fmtUs = (n) => n.toLocaleString('en-US') + 'µs';
    let onPong = null; // set by pingTest() to receive pong arrival times
    let warnedNoServerTime = false; // one-time hint when the rtt cannot be split
    let pingBaselineUs = null;      // median pong rtt, set by pingTest(); see doc comment

    // [HH:MM:SS.mmm] in browser local time, for lining logs up with the server.
    const getTimestamp = () => {
        const now = new Date();
        return now.toTimeString().split(' ')[0] + '.' + String(now.getMilliseconds()).padStart(3, '0');
    };

    const dbg = (...parts) => { if (debug) console.log(`[hx ${getTimestamp()}]`, ...parts); };
    const perfLog = (...parts) => { if (debug) console.log(`[hx-perf ${getTimestamp()}]`, ...parts); };

    // One summary line per received message, plus a late "paint" line.
    const logResponse = (id, entry, perf, bytes) => {
        if (!debug) return;

        const label = id || 'push';
        let rtt = 'n/a (untracked or already timed out)';
        if (perf.rttUs !== null) {
            rtt = fmtUs(perf.rttUs);

            if (perf.serverUs !== null) {
                // net = everything that is not the server's own work (see doc comment)
                const netUs = Math.max(0, perf.rttUs - perf.serverUs);
                rtt += ` = server ${fmtUs(perf.serverUs)} + net ${fmtUs(netUs)}`;
            } else if (pingBaselineUs !== null) {
                // Estimate: a ping's rtt is the transport cost, the rest is the server's work.
                const serverEstUs = Math.max(0, perf.rttUs - pingBaselineUs);
                rtt += ` ≈ ping-baseline ${fmtUs(pingBaselineUs)} + server-side ${fmtUs(serverEstUs)}`;
            } else if (!warnedNoServerTime) {
                warnedNoServerTime = true;
                dbg('rtt cannot be split yet: no srv=<µs> in the _hx_req_id comment and no ping baseline (run hx.pingTest())');
            }
        }

        perfLog(
            `${label} ← ${perf.outcome}`,
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

                    perfLog(
                        `${label} paint ≈ frame-wait ${fmtUs(us(tDone, tFrame))}`,
                        `+ render ${fmtUs(us(tFrame, tPainted))}`,
                        `= ${fmtUs(us(tDone, tPainted))} after handler`,
                        ...(entry ? [
                            `| felt ${fmtUs(us(entry.tStart, tPainted))} since action start`,
                            `(${fmtUs(us(entry.tStart, tPainted) - us(tDone, tFrame))} without frame-wait)`
                        ] : [])
                    );
                }, 0);
            });
        }
    };

    // Measures pure transport rtt: `count` pings, `gapMs` apart, one at a time.
    // Run hx.pingTest() in the console. Resolves with the samples in µs.
    // Compare with request rtt: flat pings + jumpy request rtt = jitter is in the
    // server's request path, not in the network or the browser.
    const pingTest = (count = 20, gapMs = 100) => new Promise((resolve) => {
        const samples = [];
        let guard = null;

        const finish = () => {
            clearTimeout(guard);
            onPong = null;

            if (samples.length > 0) {
                const sorted = [...samples].sort((x, y) => x - y);
                const medianUs = sorted[Math.floor(sorted.length / 2)];
                pingBaselineUs = medianUs; // used to estimate server-side time in logResponse

                perfLog(
                    `pingTest ${samples.length}/${count}`,
                    `| min ${fmtUs(sorted[0])}`,
                    `| median ${fmtUs(medianUs)} (now the rtt baseline)`,
                    `| max ${fmtUs(sorted[sorted.length - 1])}`
                );
            }
            resolve(samples);
        };

        const sendOne = () => {
            if (!ws || ws.readyState !== WebSocket.OPEN) {
                dbg('pingTest: socket not open');
                finish();
                return;
            }
            guard = setTimeout(finish, 2000); // pong never came back
            pingSentAt = performance.now();
            ws.send('ping');
        };

        onPong = (tRecv) => {
            clearTimeout(guard);
            samples.push(us(pingSentAt, tRecv));

            if (samples.length < count) {
                setTimeout(sendOne, gapMs);
            } else {
                finish();
            }
        };

        sendOne();
    });

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
    // Single, consistent beforeSwap / afterSwap per element.
    // requestId is optional (present for live WS responses, null for cache hits).
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

    // --- WebSocket Core ---
    function initWs() {
        if (ws?.readyState === WebSocket.CONNECTING || ws?.readyState === WebSocket.OPEN) return;

        // Always re-read the current auth token (supports token refresh)
        const authToken = getAuthToken();
        let wsUrl = WS_URL;
        if (authToken) {
            wsUrl += (wsUrl.includes('?') ? '&' : '?') + 'token=' + encodeURIComponent(authToken);
        }

        wsConnectStart = performance.now();
        try {
            ws = new WebSocket(wsUrl);
        } catch (e) {
            fireEvent('connectionError', { error: e });
            scheduleReconnect();
            return;
        }

        ws.onopen = () => {
            wsBackoff = 1000;
            wsReconnectAttempts = 0;
            dbg(`connected to ${WS_URL} in ${fmtUs(us(wsConnectStart, performance.now()))}`);
            fireEvent('connected');

            if (debug) pingTest(10, 20); // calibrates pingBaselineUs for the rtt estimate

            if (heartbeatInterval) clearInterval(heartbeatInterval);
            heartbeatInterval = setInterval(() => {
                if (ws && ws.readyState === WebSocket.OPEN) {
                    pingSentAt = performance.now();
                    ws.send('ping');
                }
            }, CONFIG.WS_HEARTBEAT_INTERVAL);
        };

        ws.onmessage = (event) => {
            const tRecv = performance.now();
            let html = event.data;

            if (html === 'pong') {
                if (onPong) {
                    onPong(tRecv); // pingTest() is collecting samples and prints its own summary
                } else if (pingSentAt) {
                    dbg(`pong, heartbeat rtt ${fmtUs(us(pingSentAt, tRecv))}`);
                }
                return;
            }

            const bytes = html.length;
            let responseId = null;
            let entry = null; // grabbed before retireRequest deletes it, needed for rtt and felt time

            const commentMatch = html.match(/^<!--\s*_hx_req_id:\s*(\S+?)(?:\s+srv=(\d+))?\s*-->\s*/);
            if (commentMatch) {
                responseId = commentMatch[1];
                entry = pendingRequests.get(responseId) || null;
            }

            // Server-reported duration (µs), optional: <!-- _hx_req_id: hx4 srv=187 -->
            const serverUs = commentMatch && commentMatch[2] !== undefined
                ? Number(commentMatch[2])
                : null;

            // rtt is fixed here, before any parse / swap / listener runs, so browser
            // work after the frame arrived can never leak into it.
            const perf = {
                tRecv,
                rttUs: entry ? us(entry.tSend, tRecv) : null,
                serverUs,
                tParseStart: tRecv,
                tParseEnd: tRecv,
                tEnd: tRecv,
                swapUs: 0,
                swapped: 0,
                outcome: 'swap'
            };

            try {
                if (commentMatch) {
                    html = html.slice(commentMatch[0].length);
                }

                if (html.startsWith('unknown endpoint:')) {
                    perf.outcome = 'unknown-endpoint';
                    retireRequest(responseId, null, { error: html });
                    fireEvent('error', { message: html, requestId: responseId });
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

                    // Fire redirect first so plugins can cancel before we retire
                    const redirectEvent = fireEvent('redirect', {
                        url: redirectUrl,
                        requestId: responseId,
                        response: html,
                        cancelable: true
                    });

                    // Always retire so pendingCount stays accurate
                    retireRequest(responseId, html, { redirected: true, redirectUrl });

                    if (!redirectEvent.defaultPrevented) {
                        window.location.href = redirectUrl;
                    }
                    return;
                }

                const elements = temp.querySelectorAll('[id]');
                if (elements.length === 0) {
                    // Empty / non-swap response – still retire cleanly
                    perf.outcome = 'no-swap';
                    retireRequest(responseId, html);
                    return;
                }

                // No batch beforeSwap – only the per-element event inside swapElementById
                const tSwap0 = performance.now();
                elements.forEach(el => swapElementById(el.id, el.outerHTML, responseId));
                perf.swapUs = us(tSwap0, performance.now());
                perf.swapped = elements.length;

                retireRequest(responseId, html);
            } finally {
                perf.tEnd = performance.now();
                logResponse(responseId, entry, perf, bytes);
            }
        };

        ws.onerror = (err) => {
            dbg('socket error', err);
            fireEvent('error', { error: err });
        };

        ws.onclose = (event) => {
            dbg(`closed code=${event.code} reason="${event.reason}" clean=${event.wasClean}; ${pendingRequests.size} request(s) in flight`);
            // Notify every in-flight request so listeners (loading indicator, etc.) can clean up
            const pendingSnapshot = Array.from(pendingRequests.entries());
            pendingRequests.clear();

            pendingSnapshot.forEach(([id, entry]) => {
                clearTimeout(entry.timerId);
                fireEvent('timeout', {
                    requestId: id,
                    element: entry.element,
                    endpoint: entry.endpoint,
                    method: entry.method,
                    reason: 'connection_closed',
                    code: event.code,
                    closeReason: event.reason
                });
                fireEvent('afterResponse', {
                    requestId: id,
                    pendingCount: 0,
                    response: null,
                    element: entry.element,
                    endpoint: entry.endpoint,
                    method: entry.method,
                    connectionClosed: true
                });
            });

            fireEvent('disconnected', { code: event.code, reason: event.reason });

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

        dbg(`reconnecting in ${Math.round(delay)}ms (attempt ${wsReconnectAttempts + 1})`);
        fireEvent('reconnecting', { delay, attempt: wsReconnectAttempts + 1 });

        wsReconnectTimer = setTimeout(() => {
            wsReconnectTimer = null;
            wsReconnectAttempts++;
            initWs();
        }, delay);
        wsBackoff = Math.min(wsBackoff * 2, CONFIG.WS_RECONNECT_MAX_BACKOFF);
    }

    /**
     * Force a clean reconnect (useful after auth-token refresh).
     * Closes the current socket (if any) and immediately starts a new connection
     * with the latest token from meta/cookie.
     */
    function reconnect() {
        if (wsReconnectTimer) {
            clearTimeout(wsReconnectTimer);
            wsReconnectTimer = null;
        }
        if (heartbeatInterval) {
            clearInterval(heartbeatInterval);
            heartbeatInterval = null;
        }
        if (ws) {
            // Avoid the normal onclose reconnect logic racing with us
            const oldWs = ws;
            ws = null;
            try {
                oldWs.onopen = null;
                oldWs.onclose = null;
                oldWs.close(1000, 'Manual reconnect');
            } catch (_) { /* ignore */ }
        }
        wsBackoff = 1000;
        wsReconnectAttempts = 0;
        initWs();
    }

    // --- Core Request Sending ---
    const sendWsAction = (element, form) => {
        const id = 'hx' + (++requestIdCounter);
        const tStart = performance.now();

        // Fire before request - plugins can cancel
        const beforeEvent = fireEvent('beforeRequest', {
            element,
            form,
            requestId: id,
            cancelable: true
        });

        const tBefore = performance.now();

        if (beforeEvent.defaultPrevented) {
            dbg(`${id} cancelled by a beforeRequest listener`);
            fireEvent('requestCancelled', {
                element,
                requestId: id
            });

            // Also fire afterResponse so UI knows to clean up
            fireEvent('afterResponse', {
                requestId: id,
                pendingCount: pendingRequests.size,
                response: null,
                element,
                endpoint: element.getAttribute(HX.GET) || element.getAttribute(HX.POST),
                method: element.hasAttribute(HX.POST) ? 'POST' : 'GET',
                cancelled: true  // ← Flag to indicate cancellation
            });

            return;// do not proceed with actual request
        }

        if (!ws || ws.readyState !== WebSocket.OPEN) {
            dbg(`${id} not sent: socket not open (readyState ${ws ? ws.readyState : 'null'})`);
            fireEvent('connectionError', { element });
            return;
        }

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

        const csrf = getCSRFToken();
        if (csrf) params.append('_csrf', csrf);

        const timerId = setTimeout(() => {
            if (pendingRequests.has(id)) {
                const entry = pendingRequests.get(id);
                pendingRequests.delete(id);
                clearTimeout(entry.timerId);
                dbg(`${id} timed out after ${fmtUs(us(entry.tSend, performance.now()))}`);

                fireEvent('timeout', {
                    requestId: id,
                    element: entry.element,
                    endpoint: entry.endpoint,
                    method: entry.method
                });

                // Always emit afterResponse so UI plugins can hide spinners etc.
                fireEvent('afterResponse', {
                    requestId: id,
                    pendingCount: pendingRequests.size,
                    response: null,
                    element: entry.element,
                    endpoint: entry.endpoint,
                    method: entry.method,
                    timedOut: true
                });
            }
        }, CONFIG.WS_REQUEST_TIMEOUT);

        pendingRequests.set(id, {
            timerId,
            tStart,
            tSend: 0, // set right before ws.send()
            element,
            endpoint,
            method: isPost ? 'POST' : 'GET'
        });
        params.append('_hx_req_id', id);

        const wire = `${isPost ? 'POST' : 'GET'} ${endpoint}\n${params.toString()}`;

        const tSend0 = performance.now();
        pendingRequests.get(id).tSend = tSend0;

        try {
            ws.send(wire);
            if (debug) {
                const tSend1 = performance.now();
                perfLog(
                    `${id} → ${isPost ? 'POST' : 'GET'} ${endpoint}`,
                    `| beforeRequest listeners ${fmtUs(us(tStart, tBefore))}`,
                    `| collect+prep ${fmtUs(us(tBefore, tSend0))}`,
                    `| ws.send() ${fmtUs(us(tSend0, tSend1))}`,
                    `| client total ${fmtUs(us(tStart, tSend1))}`,
                    `| ${wire.length} chars, buffered ${ws.bufferedAmount}B`
                );
            }
            fireEvent('requestSent', {
                requestId: id,
                element,
                endpoint,
                method: isPost ? 'POST' : 'GET'
            });
            // Emit afterRequest so cache invalidation (and other plugins) can react
            fireEvent('afterRequest', {
                requestId: id,
                element,
                endpoint,
                method: isPost ? 'POST' : 'GET'
            });
        } catch (e) {
            if (pendingRequests.has(id)) {
                clearTimeout(pendingRequests.get(id).timerId);
                pendingRequests.delete(id);
            }
            dbg(`${id} send failed`, e);
            fireEvent('error', { error: e, element });
            // Ensure afterResponse still fires for consistency
            fireEvent('afterResponse', {
                requestId: id,
                pendingCount: pendingRequests.size,
                response: null,
                element,
                endpoint,
                method: isPost ? 'POST' : 'GET',
                error: e
            });
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
                    fireEvent('redirect', { url: redirectUrl, element });
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
            sendWsAction(element, form);
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
        sendWsAction(element, form);
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
    // This ensures buttons and selects work even after DOM swaps
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
    });

    document.addEventListener('visibilitychange', () => {
        if (document.visibilityState === 'visible') {
            if (!ws || ws.readyState === WebSocket.CLOSED || ws.readyState === WebSocket.CLOSING) {
                fireEvent('visibilityRestored');
                initWs();
            }
        }
    });

    window.addEventListener('online', () => {
        fireEvent('online');
        if (!ws || ws.readyState !== WebSocket.OPEN) {
            initWs();
        }
    });

    window.addEventListener('offline', () => {
        fireEvent('offline');
    });

    // --- Expose API for plugins ---
    window.hx = {
        version: VERSION,
        fireEvent,
        sendWsAction,
        handleUpload,
        swapElementById,
        getAuthToken,
        getCSRFToken,
        reconnect,          // force reconnect with fresh token
        setDebug: (on) => { debug = Boolean(on); },
        pingTest,
        config: CONFIG,
        constants: HX
    };

    // Initialize
    fireEvent('init');
    initWs();
});