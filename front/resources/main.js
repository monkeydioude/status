let arValue = localStorage.getItem("autoReload")

class AutoReload {
    defaultValue = 60 * 1000
    timeoutId = null
    isSet = false
    inputTag = document.querySelector("[data-auto-reload]");
    containerTag = document.querySelector("[data-auto-reload-container]");

    constructor() {
        this.isSet = !!localStorage.getItem("autoReload")
        if (this.isOn()) {
            this.inputTag.checked = true;
        }
        this.containerTag.classList.remove("hidden");
        this.inputTag.addEventListener("click", (event) => {
            if (event.currentTarget.checked) {
                this.trigger();
                return;
            }
            this.cancel();
        });
    }
    getAutoReloadValue() {
        return localStorage.getItem("autoReload") || this.defaultValue;
    }
    setOn(value) {
        localStorage.setItem("autoReload", value || this.defaultValue);
        this.isSet = true;
    }
    setOff() {
        localStorage.removeItem("autoReload");
        this.isSet = false;
    }
    isOn() {
        return this.isSet;
    }
    start() {
        if (!this.isOn()) {
            return;
        }
        this.timeoutId = setTimeout(() => location.reload(), this.getAutoReloadValue());
    }
    trigger() {
        if (this.isOn() && this.timeoutId) {
            return;
        }
        this.setOn();
        this.timeoutId = setTimeout(() => location.reload(), this.getAutoReloadValue());
    }
    cancel() {
        if (this.timeoutId == null || !this.isOn())
            return;
        clearTimeout(this.timeoutId);
        this.setOff();
    }
}

window.addEventListener("load", () => {
    const autoReload = new AutoReload();
    autoReload.start();

    document
        .querySelectorAll("[data-status]")
        .forEach((elm) => {
            elm.addEventListener("click", (event) => {
                const target = event.currentTarget;
                if (target.classList.contains("full-status")) {
                    autoReload.trigger();
                    target.classList.remove("full-status");
                } else {
                    autoReload.cancel();
                    target.classList.add("full-status");
                }
            });
        });

});