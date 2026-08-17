/**
 * Caching plugin for hxgo
 * Adds client-side caching with Time-To-Live (TTL) in seconds
 * 
 * Usage:
 *   <div hx-get="/api/data" hx-cache="true" hx-cache-ttl="60">
 *   <div hx-get="/api/user" hx-cache="true" hx-cache-ttl="300">
 *   <button hx-post="/refresh" hx-cache-invalidate="/api/data">
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
            const cacheEvent = new CustomEvent('hx:cacheHit', {
                detail: {
                    cacheKey,
                    response: cachedResponse,
                    element
                }
            });
            document.dispatchEvent(cacheEvent);

            applyCachedResponse(cachedResponse, element);

            e.preventDefault();
        }
    });

    document.addEventListener('hx:afterResponse', (e) => {
        const { element } = e.detail;
        const shouldCache = element?.getAttribute('hx-cache') === 'true';
        if (!shouldCache) return;

        const ttl = parseInt(element.getAttribute('hx-cache-ttl')) || cacheManager.defaultTTL;

        // Store in cache
        // Note: Needs response data from core
    });

    document.addEventListener('hx:afterRequest', (e) => {
        const { element } = e.detail;
        const invalidateAttr = element.getAttribute('hx-cache-invalidate');

        if (invalidateAttr) {
            invalidateAttr.split(',').forEach(key => {
                cacheManager.invalidate(key.trim());
            });
        }
    });

    function buildCacheKey(method, endpoint, form) {
        const formData = new FormData(form);
        const params = new URLSearchParams(formData);
        return `${method}:${endpoint}:${params.toString()}`;
    }

    function applyCachedResponse(html, element) {
        const temp = document.createElement('div');
        temp.innerHTML = html;

        const elements = temp.querySelectorAll('[id]');
        elements.forEach(el => {
            const target = document.getElementById(el.id);
            if (target) {
                target.outerHTML = el.outerHTML;
            }
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