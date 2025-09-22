<script>
import textJson from "../assets/texts/login.json"
import {
    API_login,
    API_set_new_display_name,
    API_set_new_biography,
    API_set_new_pfp
} from "../services/user-ops"
import * as register_ops from "../services/registration";

export default {
    data: function () {
        return {
            // Variables for the forms
            username: "",
            newDisplayName: "",
            newBiography: "",
            newPfp: null,

            // Full object with user data
            userData: {
                user_id: undefined,
                display_name: undefined,
                biography: undefined,
                profile_pic: undefined
            },
            textJson: textJson,

            // Flags
            should_register: false,
            current_register_step: -1,
            can_upload_pfp: false,

            login_valid: false,
            display_name_valid: false,
            biography_valid: false,
            photo_valid: false
        }
    },

    methods: {
        // Wrapper for the login function
        async login() {
            if (!this.login_valid) {
                // TODO: Implement error message
                return
            }

            // Perform login, wait for promise
            await API_login(this.username).then((value) => {
                Object.assign(this.userData, value.userData)
                this.should_register = value.shouldRegister;
            })

            // If the user is new, perform registration
            if (this.should_register === true) {
                this.current_register_step += 1;
                await register_ops.start_registration();
                console.log(`Bro got registered with username ${this.userData.user_id} and display name ${this.userData.display_name}`);
            } else {
                // TODO: Bring to chats
                console.log(`Bro exists with username ${this.userData.user_id} and display name ${this.userData.display_name}`);
            }
        },

        // Wrapper for setting the new display name
        async set_new_display_name() {
            console.log(this.newDisplayName);
            // Perform API call, wait for promise
            await API_set_new_display_name(this.userData.user_id, this.newDisplayName).then((_) => {
                this.userData.display_name = this.newDisplayName;

                register_ops.color_new_progress_bar(this.current_register_step);
                this.current_register_step += 1;
                register_ops.start_register_new_bio();
            })
        },

        // Wrapper for setting the new biography
        async set_new_biography() {
            await API_set_new_biography(this.userData.user_id, this.newBiography).then((_) => {
                this.userData.biography = this.newBiography;

                register_ops.color_new_progress_bar(this.current_register_step);
                this.current_register_step += 1;
                register_ops.start_register_new_pfp();
            })
        },

        // Wrapper for setting the new PFP
        async set_new_pfp() {
            if (!this.photo_valid) return;

            await API_set_new_pfp(this.userData.user_id, this.userData.profile_pic).then((_) => {
                register_ops.color_new_progress_bar(this.current_register_step);
                this.current_register_step += 1;
            });
        },

        // Event to check if the image is convertible in base64
        async upload_pfp(event) {
            if (!event.target.files.length) return;

            // Save pfp and convert it to base64
            this.newPfp = event.target.files[0];
            console.log(this.newPfp);

            let reader = new FileReader();
            let base64Pfp = null;
            reader.onloadend = async () => {
                base64Pfp = reader.result
                    .replace("data:", "")
                    .replace(/^.+,/, "");

                // Check if it's valid
                console.log(base64Pfp);

                this.photo_valid = await register_ops.color_enter_button(
                    base64Pfp, 0, 4294967296, 3, this.photo_valid);

                // Use it in the user object
                if (this.photo_valid) {
                    this.userData.profile_pic = base64Pfp;
                }
            }

            reader.readAsDataURL(this.newPfp);
        },

        // Function for checking whether an input is fine or not
        async check_enter_color () {
            switch (this.current_register_step) {
                case -1:
                    console.log(this.login_valid)
                    this.login_valid = await register_ops.color_enter_button(
                        this.username, 3, 16, 0, this.login_valid);
                    break
                case 0:
                    this.display_name_valid = await register_ops.color_enter_button(
                        this.newDisplayName, 1, 32, 1, this.display_name_valid);
                    break
                case 1:
                    this.biography_valid = await register_ops.color_enter_button(
                        this.newBiography, 1, 2048, 2, this.biography_valid);
                    break
                case 2:
                    this.photo_valid = await register_ops.color_enter_button(
                        this.userData.profile_pic, 0, 4294967296, 3, this.photo_valid);
                    break
            }
        },

        // Wrapper for the library call
        async skip_registration_step() {
            await register_ops.skip_registration_step(this.current_register_step);
            this.current_register_step += 1;
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
                        <p id="register_description">Hi <b>@{{ userData.user_id }}</b>, do you want to set a new display name?
                        Don't worry, if you choose not to set it, <b>{{ userData.user_id }}</b> will be used instead.
                        You can change your display name anytime later in the settings</p>
                        <div id="login_form">
                            <input v-model="username" class="login_register_input" placeholder="Insert your username"
                                   v-on:keyup.enter="login" v-on:input="check_enter_color">
                            <button class="login_button_general login_button-unavailable" @click="login">
                                <img class="login_arrow login_arrow-unavailable"
                                     src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                            </button>
                        </div>
                        <div id="new_display_form">
                            <div class="form_zone">
                                <input v-model="newDisplayName" class="login_register_input" v-on:input="check_enter_color"
                                       placeholder="Choose a display name" v-on:keyup.enter="set_new_display_name">
                                <button class="login_button_general login_button-unavailable" @click="set_new_display_name">
                                    <img class="login_arrow login_arrow-unavailable"
                                         src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                                </button>
                            </div>
                            <button class="skip_step" @click="skip_registration_step">I don't want to set a new display name</button>
                        </div>
                        <div id="new_bio_form">
                            <div class="form_zone">
                                <input v-model="newBiography" id="set_new_bio_input" class="login_register_input"
                                       placeholder="Choose a biography" v-on:keyup.enter="set_new_biography" v-on:input="check_enter_color">
                                <button class="login_button_general login_button-unavailable" @click="set_new_biography">
                                    <img class="login_arrow login_arrow-unavailable"
                                         src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                                </button>
                            </div>
                            <button class="skip_step" @click="skip_registration_step">I don't want to set a biography</button>
                        </div>
                        <div id="new_pfp_form">
                            <div class="form_zone">
                                <label id="set_new_pfp_label">
                                    Click to upload a profile picture
                                    <input id="set_new_pfp_input" class="login_register_input" type="file"
                                           accept="image/*" @change="upload_pfp">
                                </label>
                                <button class="login_button_general login_button-unavailable"
                                        id="login_button-unavailable" @click="set_new_pfp">
                                    <img class="login_arrow login_arrow-unavailable"
                                         src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                                </button>
                            </div>
                            <button class="skip_step" @click="skip_registration_step">I don't want to set a profile picture</button>
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