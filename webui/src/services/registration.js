// The start_registration performs the animations for switching from the login
// form to the registration forms
async function start_registration() {
    // Get all elements from their IDs
    let registrationForm = document.getElementById("login_main_form");
    let loginInput = document.getElementById("login_form");
    let loginTitle = document.getElementById("login_title");
    let registrationTitle = document.getElementById("register_title");
    let registrationDescription = document.getElementById("register_description");
    let newDisplayNameForm = document.getElementById("new_display_form");
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
        newDisplayNameForm.style.display = "flex";
        progressBars.style.display = "flex";
    }, 800);

    setTimeout(function() {
        registrationTitle.style.transition = "margin-top 1s ease-in-out";
        registrationTitle.style.marginTop = "-10px";
    }, 2000);

    setTimeout(function() {
        registrationDescription.style.transition = "opacity 1s ease-in-out";
        registrationDescription.style.opacity = "100%";

        newDisplayNameForm.style.transition = "opacity 1s ease-in-out";
        newDisplayNameForm.style.opacity = "1";

        progressBars.style.transition = "opacity 1s ease-in-out";
        progressBars.style.opacity = "1";
    }, 3100);
}

async function start_register_new_bio() {
    let registrationDescription = document.getElementById("register_description");
    let newDisplayNameForm = document.getElementById("new_display_form");
    let newBioForm = document.getElementById("new_bio_form");

    newDisplayNameForm.style.transition = "opacity 1s ease-in-out";
    registrationDescription.style.transition = "opacity 1s ease-in-out";

    newDisplayNameForm.style.opacity = "0";
    registrationDescription.style.opacity = "0";

    setTimeout(function() {
        newDisplayNameForm.style.display = "none";
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

async function start_register_new_pfp() {
    let registrationDescription = document.getElementById("register_description");
    let newBioForm = document.getElementById("new_bio_form");
    let newPfpForm = document.getElementById("new_pfp_form");
    // let confirmButton = document.getElementById("login_button");

    newBioForm.style.transition = "opacity 1s ease-in-out";
    registrationDescription.style.transition = "opacity 1s ease-in-out";

    newBioForm.style.opacity = "0";
    registrationDescription.style.opacity = "0";

    setTimeout(function() {
        newBioForm.style.display = "none";
        newPfpForm.style.display = "flex";

        // inputForm.reset()
        registrationDescription.innerHTML = `You can set up also a profile picture. In the case where you don't want to
        upload a picture of your own, we'll provide an automatic one for you. You can also change it later in 
        the settings`
    }, 1010);

    setTimeout(function() {
        newPfpForm.style.transition = "opacity 1s ease-in-out";
        newPfpForm.style.opacity = "1";
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
            await start_register_new_pfp();
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

// check_enter_button changes the color and the flags
async function color_enter_button(val_to_check, minLen, maxLen, element, flag_to_check) {
    // Select button and arrow
    let button = document.getElementsByClassName("login_button_general")[element];
    let arrow = document.getElementsByClassName("login_arrow")[element];

    // Change the arrow and button
    if (val_to_check.length >= minLen && val_to_check.length <= maxLen && !flag_to_check) {
        arrow.classList.replace("login_arrow-unavailable", "login-arrow-available")
        button.classList.replace("login_button-unavailable", "login_button-available");

        return true;
    } else if ((val_to_check.length <= minLen || val_to_check.length >= maxLen) && flag_to_check)  {
        arrow.classList.replace("login_arrow-available", "login_arrow-unavailable");
        button.classList.replace("login_button-available", "login_button-unavailable");

        return false;
    }

    return flag_to_check;
}

export {
    start_registration,
    start_register_new_bio,
    start_register_new_pfp,

    skip_registration_step,
    color_new_progress_bar,

    color_enter_button
};