<script>
import text_json from "../assets/texts/login.json"
// import prepare_set_new_display_name from "../services/registration"

export default {
    data: function () {
        return {
            username: "",
            newDisplayName: "",

            user_data: {},
            text_json: text_json,

            current_register_step: 0
        }
    },

    methods: {
        async login() {
            try {
                // Get from the backend the user. It will be created in case it doesn't exist
                let response = await this.$axios.post("/session", {
                    headers: {
                        "Content-Type": "application/json",
                    },
                    name: this.username
                });

                this.user_data = response.data;
                let shouldRegister = response.status === 201;

                if (shouldRegister === true) {
                    await this.start_registration();
                }

                // Here, switch view
                console.log(this.user_data);
                console.log(response.status, shouldRegister);
            } catch (error) {
                console.log(error);
            }
        },

        async start_registration() {
            let registrationForm = document.getElementById("login_main_form");
            registrationForm.id = "register_main_form";

            let loginInput = document.getElementById("login_form");
            let loginTitle = document.getElementById("login_title");
            let registrationTitle = document.getElementById("register_title");
            let registrationDescription = document.getElementById("register_description");
            let newDisplayForm = document.getElementById("new_display_form");
            let progressBars = document.getElementById("progress_bars");

            await this.prepare_set_new_display_name(loginInput, loginTitle, registrationTitle,
                                                registrationDescription, newDisplayForm, progressBars);
        },

        async prepare_set_new_display_name(loginInput, loginTitle, registrationTitle,
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
        },

        async set_new_display_name() {
            try {
                let response = await this.$axios.post(`/user/${this.username}/name`,
                    {
                        new_display_name: this.newDisplayName
                    },
                    {
                        headers: {
                            Authorization: `Bearer ${this.username}`,
                            "Content-Type": "application/json",
                        },
                    }
                );

                console.log(response.data)
                await this.color_new_progress_bar()

                this.current_register_step += 1;
            } catch (e) {
                console.log("Error!")
            }
        },

        async skip_register_step() {
            switch (this.current_register_step) {
                case 0:
                    console.log("Test for skip");
                    await this.color_new_progress_bar()
                    break
                case 1:
                    // Go to setPhoto
                default:
                    // Go to WASAText
            }
        },

        async color_new_progress_bar() {
            let progBar;
            switch (this.current_register_step) {
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

            progBar.style.transition = "background-color 0.4s ease-in";
            progBar.style.backgroundColor = "var(--progress-done-color)";
        }
    }
}
</script>

<template>
    <div id="login_bg_1">
        <div id="login_bg_2">
            <div id="login_bg_3">
                <div id="login_body">
                    <div id="login_main_form">
                        <h1 id="login_title">WASAText</h1>
                        <h1 id="register_title">Welcome to WASAText</h1>
                        <p id="register_description">Hi&nbsp<b>@{{ user_data.user_id }}</b>, do you want to set a new display name?</p>
                        <div id="login_form">
                            <input v-model="username" class="login_register_input" placeholder="Insert your username" v-on:keyup.enter="login">
                            <button id="login_button" @click="login">
                                <img id="login_arrow" src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                            </button>
                        </div>
                        <div id="new_display_form">
                            <div class="form_zone">
                                <input v-model="newDisplayName" class="login_register_input" placeholder="Choose a display name" v-on:keyup.enter="set_new_display_name">
                                <button id="login_button" @click="set_new_display_name">
                                    <img id="login_arrow" src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                                </button>
                            </div>
                            <button class="skip_step" @click="skip_register_step">I don't want to set a new display name</button>
                        </div>
                        <div id="progress_bars">
                            <div class="progress" id="register_dispname"></div>
                            <div class="progress" id="register_bio"></div>
                            <div class="progress" id="register_photo"></div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>

</style>