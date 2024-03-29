const activeToasts = [];

function showToast(event) {
    if (event.detail.successful && event.target.parentElement.tagName == "DIALOG") {
        const modal = event.target.parentElement;
        if (modal) {
            modal.close();
            const form = modal.getElementsByTagName('form')[0];
            if (form) {
                form.reset();
            }
        }
    }

    createToast(event.detail.xhr.response, event.detail.successful);
}

function createToast(message, success) {
    const container = document.getElementById("toast-container");
    if (!container) return;

    const toast = document.createElement("div");
    toast.classList.add("toast");
    if (success) {
        toast.classList.add("success");
    } else {
        toast.classList.add("error");
    }
    toast.innerHTML = message;

    if (activeToasts.length > 0) {
        const previousToastRect = activeToasts[activeToasts.length - 1].getBoundingClientRect();
        toast.style.top = `${previousToastRect.top + previousToastRect.height + 8}px`;
    }

    container.appendChild(toast);
    activeToasts.push(toast);

    setTimeout(() => {
        toast.style.opacity = 0;
    }, 3000);

    setTimeout(() => {
        removeToast(toast);
    }, 4000);
}

function removeToast(toast) {
    const toastIdx = activeToasts.indexOf(toast);
    if (toastIdx !== -1) {
        activeToasts.splice(toastIdx, 1);
    }

    toast.remove();

    if (activeToasts.length == 0) return;

    activeToasts[0].style.top = '8px';
    for (let i = 1; i < activeToasts.length; i++) {
        const previousToastRect = activeToasts[i - 1].getBoundingClientRect();
        activeToasts[i].style.top = `${previousToastRect.top + previousToastRect.height + 8}px`;
    }
}
