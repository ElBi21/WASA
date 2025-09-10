async function prepare_set_new_display_name(loginInput, loginTitle, registrationTitle,
    registrationDescription, newDisplayForm, progressBars) {
    loginInput.style.transition = "opacity 0.4s ease-in-out";
    loginInput.style.opacity = "0%";

    loginTitle.style.transition = "opacity 0.4s ease-in-out";
    loginTitle.style.opacity = "0%";

    setTimeout(function() {
        loginInput.style.display = "none";
        loginTitle.style.display = "none";
    }, 400);

    registrationTitle.style.display = "flex";

    setTimeout(function() {
        registrationTitle.style.transition = "opacity 1s ease-in-out";
        registrationTitle.style.opacity = "100%";

        registrationDescription.style.display = "flex";
        newDisplayForm.style.display = "flex";
        progressBars.style.display = "flex";
    }, 800);

    setTimeout(function() {
        registrationTitle.style.transition = "margin-top 1s ease-in-out";
        registrationTitle.style.marginTop = "-10px";
    }, 2000);

    setTimeout(function() {
        //registrationTitle.style.marginBottom = "20px";
        registrationDescription.style.transition = "opacity 1s ease-in-out";
        registrationDescription.style.opacity = "100%";

        newDisplayForm.style.transition = "opacity 1s ease-in-out";
        newDisplayForm.style.opacity = "1";
        // newDisplayForm.style.marginTop = "80px";

        progressBars.style.transition = "opacity 1s ease-in-out";
        progressBars.style.opacity = "1";
        // progressBars.style.marginBottom = "-50px";
    }, 3100);

    setTimeout(function() {
    }, 2200);

    return void 0;
}

export default prepare_set_new_display_name()