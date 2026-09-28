# hxgo Cache Plugin - Complete Documentation

## Overview

The cache plugin adds client-side caching with Time-To-Live (TTL) to hxgo. Responses are stored in memory and served instantly for repeat requests within the TTL period.

## Installation

Add the cache plugin script after the core:

```html
<script src="/hxgo_core_ws.js"></script>
<script src="/hxgo_plugin_cache.js"></script>
```

## Basic Usage

### Simple Caching

Cache a response for 60 seconds (default TTL if not specified):

```html
<div hx-get="/api/user-profile" hx-cache="true">
    Loading...
</div>
```

### Cache with Specific TTL

Cache for 300 seconds (5 minutes):

```html
<div hx-get="/api/notifications" hx-cache="true" hx-cache-ttl="300">
    Loading...
</div>
```

Cache for 1 Hour

```html
<div hx-get="/api/static-content" hx-cache="true" hx-cache-ttl="3600">
    Loading...
</div>
```

### Cache Invalidation

#### Invalidate Single Cache

```html
<button hx-post="/api/update-profile" hx-cache-invalidate="/api/user-profile">
    Update Profile
</button>
```

#### Invalidate Multiple Caches

```html
<button hx-post="/api/logout" hx-cache-invalidate="/api/user-profile, /api/notifications">
    Logout
</button>
```

### Advanced Usage

#### Programmatic Cache Control

Access the cache API through window.hx.cache:

```js
// Get from cache
const cached = window.hx.cache.get('/api/user-profile');

// Set cache manually
window.hx.cache.set('/api/user-profile', '<div id="profile">John</div>', 300);

// Invalidate specific key
window.hx.cache.invalidate('/api/user-profile');

// Invalidate by prefix
window.hx.cache.invalidatePrefix('/api/user');

// Clear entire cache
window.hx.cache.clear();
```

### Cache Events

Listen for cache-related events:

```js
// Cache hit - response served from cache
document.addEventListener('hx:cacheHit', (e) => {
    console.log('Served from cache:', e.detail.cacheKey);
});

// Cache miss - fetching from server
document.addEventListener('hx:cacheMiss', (e) => {
    console.log('Cache miss, fetching from server');
});

// Cache set - response stored in cache
document.addEventListener('hx:cacheSet', (e) => {
    console.log('Cached response for:', e.detail.cacheKey);
});
```

### Use Cases

#### User Profile (5 minutes cache)

```html
<div hx-get="/api/profile" hx-cache="true" hx-cache-ttl="300">
    Loading profile...
</div>

<button hx-post="/api/update-profile" hx-cache-invalidate="/api/profile">
    Edit Profile
</button>
```

#### Navigation Menu (1 hour cache)

```html
<nav hx-get="/api/menu" hx-cache="true" hx-cache-ttl="3600">
    Loading menu...
</nav>
```

#### Notifications (30 seconds cache)

```html
<div hx-get="/api/notifications" hx-cache="true" hx-cache-ttl="30">
    Loading notifications...
</div>
```

#### Static Content (1 day cache)

```html
<div hx-get="/api/terms-of-service" hx-cache="true" hx-cache-ttl="86400">
    Loading...
</div>
```

## Debugging

Enable console logging to debug cache behavior:

```js
document.addEventListener('hx:cacheHit', (e) => {
    console.log('CACHE HIT:', e.detail.cacheKey);
});

document.addEventListener('hx:cacheMiss', (e) => {
    console.log('CACHE MISS:', e.detail.cacheKey);
});

// Inspect cache
console.log(window.hx.cache);
```

## Example: Complete Page

```html
<!DOCTYPE html>
<html>
<head>
    <title>Cache Example</title>
    <script src="/hxgo_core_ws.js"></script>
    <script src="/hxgo_plugin_cache.js"></script>
</head>
<body>
    <!-- Cached for 5 minutes -->
    <div id="user-profile" hx-get="/api/profile" hx-cache="true" hx-cache-ttl="300">
        Loading profile...
    </div>

    <!-- Cached for 30 seconds -->
    <div id="notifications" hx-get="/api/notifications" hx-cache="true" hx-cache-ttl="30">
        Loading notifications...
    </div>

    <!-- Invalidate profile cache on update -->
    <button hx-post="/api/update-profile" hx-cache-invalidate="/api/profile">
        Update Profile
    </button>

    <!-- Clear all user caches on logout -->
    <button hx-post="/api/logout" hx-cache-invalidate="/api/profile, /api/notifications">
        Logout
    </button>

    <script>
        // Debug cache
        document.addEventListener('hx:cacheHit', (e) => {
            console.log('Cache hit:', e.detail.cacheKey);
        });
    </script>
</body>
</html>
```

## Troubleshooting

Cache not working?

- Check that hx-cache="true" is set
- Verify hx-cache-ttl is a valid number (seconds)
- Make sure endpoint is correct
- Check browser console for errors

Stale data?

- Reduce TTL value
- Ensure proper cache invalidation
- Use hx-cache-invalidate after updates

Memory issues?

- Clear cache periodically: window.hx.cache.clear()
- Reduce TTL values
- Limit number of cached endpoints

