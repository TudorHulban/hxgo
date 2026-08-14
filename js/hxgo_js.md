# hxgo - WebSocket-First HTML Over the Wire Framework

A lightweight, event-driven JavaScript framework for building dynamic web applications with WebSocket communication. hxgo follows the htmx philosophy of declarative HTML attributes but is designed specifically for WebSocket-first architectures.

**Version:** `0.1.1-core`

## Overview

hxgo allows you to build interactive web applications by adding declarative attributes to your HTML. The framework handles WebSocket communication, DOM swapping, form validation, and file uploads through a modular, event-driven architecture.

### Simple Example

```html
<button hx-post="/api/action" hx-disable="#btn">Click Me</button>
<div id="result"></div>
```

## Architecture

hxgo is built around a minimal core with plugin modules:

```text
  ┌─────────────────────────────────────────┐
  │           hxgo Core (hxgo_core_ws.js)   │
  │  - WebSocket connection management      │
  │  - Event system                         │
  │  - Request/response handling            │
  │  - DOM swapping                         │
  │  - File upload (XHR)                    │
  └─────────────────┬───────────────────────┘
                    │ Events
        ┌───────────┼───────────┐
        │           │           │
        ▼           ▼           ▼
  ┌──────────┐ ┌─────────┐ ┌──────────────┐
  │    UI    │ │Validation││   Listeners  │
  │  Plugin  │ │  Plugin  ││    Plugin    │
  └──────────┘ └─────────┘ └──────────────┘
```

## Core Module (`hxgo_core_ws.js`)

The minimal foundation that handles:

- WebSocket lifecycle (connect, reconnect, heartbeat)
- Custom event system
- Request correlation (request/response matching)
- DOM element swapping
- File uploads via XHR
- Event delegation for dynamic elements

## Plugins

### UI Plugin (`hxgo_plugin_ui.js`)

User interface concerns:

- Loading indicators
- Error alerts/popovers
- Button disabling/enabling
- Visual feedback

### Validation Plugin (`hxgo_plugin_validation.js`)

Form validation:

- Required field checking
- Min/max length validation
- Password matching
- Target element disabling

### Listeners Plugin (`hxgo_plugin_listeners.js`)

Event listener management:

- Reattaches event handlers after DOM swaps
- Show/hide element handling
- Search clearing functionality

## HTML Attributes Reference

### HTTP Methods

- `hx-get`: Send GET request over WebSocket
- `hx-post`: Send POST request over WebSocket
- `hx-upload`: Upload file via HTTP

### Element Control

- `hx-enable`: Enable elements after request
- `hx-disable`: Disable elements during request
- `hx-swap`: Target elements for content swap
- `hx-send`: Include additional elements in request
- `hx-redirect`: Redirect to URL after request

### Validation

- `hx-require`: Comma-separated required field IDs
- `hx-min`: Minimum length validation
- `hx-max`: Maximum length validation
- `hx-vdisable`: Disable target if validation fails
- `hx-pdisable`: Password match validation group

### Visibility

- `hx-show`: Show target element
- `hx-hide`: Hide target element
- `hx-show-onload`: Show element on page load

## Events

hxgo fires custom events that plugins and user code can listen to:

### Request Lifecycle Events

- `hx:beforeRequest`: Before sending request (cancelable)
- `hx:requestSent`: Request sent successfully
- `hx:afterResponse`: Response received
- `hx:requestCancelled`: Request was cancelled
- `hx:timeout`: Request timed out

### Connection Events

- `hx:connected`: WebSocket connected
- `hx:disconnected`: WebSocket disconnected
- `hx:reconnecting`: Attempting to reconnect
- `hx:connectionError`: Connection error occurred

### DOM Events

- `hx:beforeSwap`: Before DOM element swap
- `hx:afterSwap`: After DOM element swap
- `hx:reattachListeners`: Request to reattach listeners

### Upload Events

- `hx:beforeUpload`: Before file upload (cancelable)
- `hx:uploadStarted`: Upload started
- `hx:uploadProgress`: Upload progress update
- `hx:uploadComplete`: Upload completed
- `hx:uploadCancelled`: Upload was cancelled

### Other Events

- `hx:init`: Framework initialized
- `hx:redirect`: Redirect requested
- `hx:error`: Error occurred
- `hx:beforeUnload`: Page unloading

## Examples

### Basic Form Submission

```html
<form id="loginForm" hx-post="/login">
    <input type="text" name="username" hx-require="#username" hx-min="3">
    <input type="password" name="password" hx-require="#password" hx-min="6">
    <button type="submit" hx-post="/login" id="loginButton">Login</button>
</form>
```

### Password Validation

```html
<input type="password" id="password" hx-pdisable="#password, #confirmPassword, #registerBtn">
<input type="password" id="confirmPassword" hx-pdisable="#password, #confirmPassword, #registerBtn">
<button id="registerBtn" disabled>Register</button>
```

### File Upload

```html
<button hx-upload="/upload" hx-swap="#fileList" hx-redirect="/dashboard">
    Upload File
</button>
<div id="fileList"></div>
```

### Show/Hide Elements

```html
<button hx-show="#notification:3000:fade-in">Show Notification</button>
<button hx-hide="#notification">Hide Notification</button>
<div id="notification" style="display:none">Hello!</div>
```

## Extending hxgo

### Creating a Custom Plugin

```javascript
// my-analytics-plugin.js
(function() {
    // Wait for hxgo to initialize
    document.addEventListener('hx:init', () => {
        console.log('Analytics plugin initialized');
    });

    // Track all requests
    document.addEventListener('hx:beforeRequest', (e) => {
        const { element, form } = e.detail;
        const endpoint = element.getAttribute('hx-post') || element.getAttribute('hx-get');
        
        // Send to analytics
        analytics.track('request', {
            endpoint,
            timestamp: Date.now()
        });
    });

    // Track errors
    document.addEventListener('hx:error', (e) => {
        analytics.track('error', e.detail);
    });
})();
```

### Adding Custom Event Handler

```javascript
// Custom loading indicator
document.addEventListener('hx:beforeRequest', () => {
    document.body.classList.add('loading');
});

document.addEventListener('hx:afterResponse', (e) => {
    if (e.detail.pendingCount === 0) {
        document.body.classList.remove('loading');
    }
});
```

### Cancelling Requests

```javascript
// Block requests to certain endpoints
document.addEventListener('hx:beforeRequest', (e) => {
    const endpoint = e.detail.element.getAttribute('hx-post');
    if (endpoint === '/delete-account') {
        if (!confirm('Are you sure?')) {
            e.preventDefault(); // Cancel the request
        }
    }
});
```

## Backend Integration

### WebSocket Protocol

Requests are sent in this format:

```http
POST /endpoint
param1=value1&param2=value2&_hx_req_id=hx1
```

### Response Format

Responses should include the request ID as an HTML comment:

```html
<!-- _hx_req_id: hx1 -->
<div id="element1">Updated content</div>
```

### Redirect Response

```html
<!-- _hx_req_id: hx1 -->
<div id="redirect" hx-redirect="/dashboard"></div>
```

## Installation

1. Include scripts in order:

   ```html
   <script src="/hxgo_core_ws.js"></script>
   <script src="/hxgo_plugin_ui.js"></script>
   <script src="/hxgo_plugin_validation.js"></script>
   <script src="/hxgo_plugin_listeners.js"></script>
   ```

2. Add `hx-*` attributes to your HTML.
3. Configure your backend WebSocket endpoint at `/ws`.

## Configuration

Configuration is in `hxgo_core_ws.js`:

```javascript
const CONFIG = Object.freeze({
    WS_REQUEST_TIMEOUT: 30000,          // Request timeout (ms)
    WS_HEARTBEAT_INTERVAL: 30000,       // Heartbeat interval (ms)
    WS_RECONNECT_MAX_BACKOFF: 30000,    // Max reconnect delay (ms)
    WS_RECONNECT_JITTER_MAX: 5000,      // Max jitter (ms)
    MS_FILE_UPLOAD_RESET_TIMEOUT: 1500  // File input reset (ms)
});
```

## Browser Support

Modern browsers with WebSocket support:

- Chrome 16+
- Firefox 11+
- Safari 6+
- Edge 12+

## License

Proprietary / All Rights Reserved