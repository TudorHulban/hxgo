document.addEventListener('DOMContentLoaded', () => {
    const CONFIG = {
        MS_LOADING_INDICATOR_DELAY: 100,
        MS_DISABLE_TRIGGER_BUTTON: 500,
        MS_DURATION_DISPLAY_POPOVER: 3000,
        MS_DURATION_DISPLAY_ERROR_ALERT: 7000
    };

    let loadingIndicatorTimeout;
    
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
            if (popover.parentNode) document.body.removeChild(popover);
        }, CONFIG.MS_DURATION_DISPLAY_POPOVER);
    }

    function showErrorAlert(message) {
        const errorAlertBox = document.createElement('div');
        errorAlertBox.className = 'error-alert';
        errorAlertBox.textContent = message;
        document.body.appendChild(errorAlertBox);

        setTimeout(() => {
            if (errorAlertBox.parentNode) document.body.removeChild(errorAlertBox);
        }, CONFIG.MS_DURATION_DISPLAY_ERROR_ALERT);
    }

    // Loading indicator events
    document.addEventListener('hx:beforeRequest', () => showLoadingIndicator());
    document.addEventListener('hx:afterResponse', (e) => {
        if (e.detail.pendingCount === 0) hideLoadingIndicator();
    });
    document.addEventListener('hx:requestCancelled', hideLoadingIndicator);
    document.addEventListener('hx:timeout', hideLoadingIndicator);
    document.addEventListener('hx:connectionError', hideLoadingIndicator);

    document.addEventListener('hx:beforeUpload', () => showLoadingIndicator());
    document.addEventListener('hx:uploadComplete', hideLoadingIndicator);
    document.addEventListener('hx:uploadCancelled', hideLoadingIndicator);

    // Button disabling
    document.addEventListener('hx:beforeRequest', (e) => {
        const { element } = e.detail;
        element.disabled = true;
        setTimeout(() => { element.disabled = false; }, CONFIG.MS_DISABLE_TRIGGER_BUTTON);
    });

    // Error handling
    document.addEventListener('hx:error', (e) => {
        if (e.detail.message) showErrorAlert(e.detail.message);
        else if (e.detail.error) showErrorAlert('Error: ' + e.detail.error.message);
    });

    document.addEventListener('hx:timeout', () => {
        showErrorAlert('Request timed out');
    });

    document.addEventListener('hx:connectionError', () => {
        showErrorAlert('Not connected to server. Reconnecting...');
    });

    // Expose UI functions
    window.hx.ui = {
        showPopover,
        showErrorAlert,
        showLoadingIndicator,
        hideLoadingIndicator
    };
});