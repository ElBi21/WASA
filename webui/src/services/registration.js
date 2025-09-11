// The start_registration performs the animations for switching from the login
// form to the registration forms
async function start_registration() {
    // Get all elements from their IDs
    let registrationForm = document.getElementById("login_main_form");
    let loginInput = document.getElementById("login_form");
    let loginTitle = document.getElementById("login_title");
    let registrationTitle = document.getElementById("register_title");
    let registrationDescription = document.getElementById("register_description");
    let newDisplayForm = document.getElementById("new_display_form");
    let progressBars = document.getElementById("progress_bars");

    // Start performing all the animations
    registrationForm.id = "register_main_form";
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

        registrationDescription.style.display = "inline";
        newDisplayForm.style.display = "flex";
        progressBars.style.display = "flex";
    }, 800);

    setTimeout(function() {
        registrationTitle.style.transition = "margin-top 1s ease-in-out";
        registrationTitle.style.marginTop = "-10px";
    }, 2000);

    setTimeout(function() {
        registrationDescription.style.transition = "opacity 1s ease-in-out";
        registrationDescription.style.opacity = "100%";

        newDisplayForm.style.transition = "opacity 1s ease-in-out";
        newDisplayForm.style.opacity = "1";

        progressBars.style.transition = "opacity 1s ease-in-out";
        progressBars.style.opacity = "1";
    }, 3100);
}

async function start_register_new_bio() {
    let newDispNameForm = document.getElementById("new_display_form")
    let registrationDescription = document.getElementById("register_description");
    let newBioForm = document.getElementById("new_bio_form")
    // let inputForm = document.getElementById("set_new_bio_input")

    newDispNameForm.style.transition = "opacity 1s ease-in-out";
    registrationDescription.style.transition = "opacity 1s ease-in-out";

    newDispNameForm.style.opacity = "0";
    registrationDescription.style.opacity = "0";

    setTimeout(function() {
        newDispNameForm.style.display = "none";
        newBioForm.style.display = "flex";
        // inputForm.reset()
        registrationDescription.innerHTML = `Every user can set its own biography; if you want, you can set one now.
                In case you don't want to set it up now, don't worry, you'll be able to change it in a later moment.`
    }, 1010);

    setTimeout(function() {
        newBioForm.style.transition = "opacity 1s ease-in-out";
        newBioForm.style.opacity = "1";
        registrationDescription.style.opacity = "1";
    }, 1050);
}


async function skip_registration_step(register_step) {
    switch (register_step) {
        case 0:
            console.log("Skipping display name");
            await start_register_new_bio();
            break
        case 1:
            // Go to setPhoto
            console.log("Skipping bio");
            break
        default:
            // Go to WASAText
    }

    await color_new_progress_bar(register_step);
}

async function color_new_progress_bar(register_step) {
    let progBar;

    // Select progress bar to color
    switch (register_step) {
        case 0:
            progBar = document.getElementById("register_dispname")
            break
        case 1:
            progBar = document.getElementById("register_bio")
            break
        default:
            progBar = document.getElementById("register_photo")
            break
    }

    // Perform animation
    progBar.style.transition = "background-color 0.4s ease-in";
    progBar.style.backgroundColor = "var(--progress-done-color)";
}

export {
    start_registration,
    start_register_new_bio,

    skip_registration_step,
    color_new_progress_bar
};