/**
 * Caching plugin for hxgo
 * Adds client-side caching with Time-To-Live (TTL) in seconds
 * 
 * Usage:
 *   <div hx-get="/api/data" hx-cache="true" hx-cache-ttl="60">
 *   <div hx-get="/api/user" hx-cache="true" hx-cache-ttl="300">
 *   <button hx-post="/refresh" hx-cache-invalidate="/api/data">
 *   <button hx-post="/refresh" hx-cache-invalidate="/api/data,/api/user">
 */

(function () {
    class CacheManager {
        constructor(maxEntries = 100, defaultTTL = 300) {
            this.cache = new Map();
            this.maxEntries = maxEntries;
            this.defaultTTL = defaultTTL;
        }

        get(key) {
            const entry = this.cache.get(key);
            if (!entry) return null;

            if (Date.now() > entry.expiresAt) {
                this.cache.delete(key);
                return null;
            }

            return entry.value;
        }

        set(key, value, ttlSeconds = this.defaultTTL) {
            if (this.cache.size >= this.maxEntries) {
                const oldestKey = this.cache.keys().next().value;
                this.cache.delete(oldestKey);
            }

            this.cache.set(key, {
                value,
                expiresAt: Date.now() + (ttlSeconds * 1000)
            });
        }

        invalidate(key) {
            this.cache.delete(key);
        }

        invalidatePrefix(prefix) {
            for (const key of this.cache.keys()) {
                if (key.startsWith(prefix)) {
                    this.cache.delete(key);
                }
            }
        }

        clear() {
            this.cache.clear();
        }
    }

    const cacheManager = new CacheManager(
        window.hx.config.CACHE_MAX_ENTRIES,
        window.hx.config.CACHE_DEFAULT_TTL
    );

    document.addEventListener('hx:beforeRequest', (e) => {
        const { element, form } = e.detail;

        const shouldCache = element.getAttribute('hx-cache') === 'true';
        if (!shouldCache) return;

        const endpoint = element.getAttribute('hx-get') || element.getAttribute('hx-post');
        const method = element.hasAttribute('hx-post') ? 'POST' : 'GET';

        const cacheKey = buildCacheKey(method, endpoint, form);

        const cachedResponse = cacheManager.get(cacheKey);
        if (cachedResponse) {
            // Fire a proper cacheHit event
            window.hx.fireEvent('cacheHit', {
                cacheKey,
                response: cachedResponse,
                element,
                endpoint,
                method
            });

            // Apply using the same swap path as the core
            applyCachedResponse(cachedResponse);

            // Synthesize afterResponse so UI plugin hides the loading indicator
            // and any other afterResponse listeners still run
            window.hx.fireEvent('afterResponse', {
                requestId: null,
                pendingCount: 0,          // critical for loading indicator
                response: cachedResponse,
                element,
                endpoint,
                method,
                fromCache: true
            });

            e.preventDefault();
        }
    });

    document.addEventListener('hx:afterResponse', (e) => {
        // Skip if this is the synthetic event we just fired ourselves
        if (e.detail.fromCache) return;

        const { element, response, endpoint, method } = e.detail;
        const shouldCache = element?.getAttribute('hx-cache') === 'true';
        if (!shouldCache || !response) return;

        const ttl = parseInt(element.getAttribute('hx-cache-ttl')) || cacheManager.defaultTTL;
        const form = element.closest('form') || document.createElement('form');
        const cacheKey = buildCacheKey(
            method || 'GET',
            endpoint || element.getAttribute('hx-get') || element.getAttribute('hx-post'),
            form
        );

        cacheManager.set(cacheKey, response, ttl);
    });

    document.addEventListener('hx:afterRequest', (e) => {
        const { element } = e.detail;
        const invalidateAttr = element.getAttribute('hx-cache-invalidate');

        if (!invalidateAttr) return;

        invalidateAttr.split(',').forEach(raw => {
            const key = raw.trim();
            if (!key) return;

            // Exact match (in case someone puts a full key)
            cacheManager.invalidate(key);

            // Endpoint-style invalidation (the common case)
            ['GET', 'POST'].forEach(method => {
                // Match both "GET:/api/data" and "GET:/api/data:params..."
                cacheManager.invalidatePrefix(`${method}:${key}`);
            });
        });
    });

    function buildCacheKey(method, endpoint, form) {
        const formData = new FormData(form);
        const params = new URLSearchParams(formData);
        return `${method}:${endpoint}:${params.toString()}`;
    }

    function applyCachedResponse(html) {
        const temp = document.createElement('div');
        temp.innerHTML = html;

        const redirectElement = temp.querySelector('[hx-redirect]');
        if (redirectElement) {
            const redirectUrl = redirectElement.getAttribute('hx-redirect');
            const redirectEvent = window.hx.fireEvent('redirect', {
                url: redirectUrl,
                fromCache: true,
                cancelable: true
            });
            if (!redirectEvent.defaultPrevented) {
                window.location.href = redirectUrl;
            }
            return;
        }

        const elements = temp.querySelectorAll('[id]');
        if (elements.length === 0) return;

        // No batch beforeSwap – only the per-element event inside swapElementById
        // (requestId = null for cache hits)
        elements.forEach(el => {
            window.hx.swapElementById(el.id, el.outerHTML, null);
        });
    }

    window.hx.cache = {
        get: (key) => cacheManager.get(key),
        set: (key, value, ttl) => cacheManager.set(key, value, ttl),
        invalidate: (key) => cacheManager.invalidate(key),
        invalidatePrefix: (prefix) => cacheManager.invalidatePrefix(prefix),
        clear: () => cacheManager.clear()
    };
})();