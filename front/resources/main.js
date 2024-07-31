window.addEventListener("load", () => {
    setTimeout(() => location.reload(), 60000)
    document
        .querySelectorAll("[data-status]")
        .forEach((elm) => {
            elm.addEventListener("click", (event) => {
                const target = event.currentTarget;
                if (target.classList.contains("full-status")) {
                    target.classList.remove("full-status")
                } else {
                    target.classList.add("full-status")
                }
            });
        })

});