<script>
import textJson from "../assets/texts/login.json"
import {
    API_login,
    API_set_new_display_name,
    API_set_new_biography
} from "../services/user-ops"
import * as register_ops from "../services/registration";

export default {
    data: function () {
        return {
            // Variables for the forms
            username: "",
            newDisplayName: "",
            newBiography: "",

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
            current_register_step: 0
        }
    },

    methods: {
        async login() {
            // Perform login, wait for promise
            await API_login(this.username).then((value) => {
                Object.assign(this.userData, value.userData)
                this.should_register = value.shouldRegister;
            })

            // If the user is new, perform registration
            if (this.should_register === true) {
                await register_ops.start_registration();
                console.log(`Bro got registered with username ${this.userData.user_id} and display name ${this.userData.display_name}`);
            } else {
                // TODO: Bring to chats
                console.log(`Bro exists with username ${this.userData.user_id} and display name ${this.userData.display_name}`);
            }
        },

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

        async set_new_biography() {
            await API_set_new_biography(this.userData.user_id, this.newBiography).then((_) => {
                this.userData.biography = this.newBiography;

                register_ops.color_new_progress_bar(this.current_register_step);
                this.current_register_step += 1;
                // Start register new photo
                // regist_ops.start_register_new_bio();
            })
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
                            <button class="skip_step" @click="skip_registration_step">I don't want to set a new display name</button>
                        </div>
                        <div id="new_bio_form">
                            <div class="form_zone">
                                <input v-model="newBiography" id="set_new_bio_input" class="login_register_input" placeholder="Choose a biography" v-on:keyup.enter="set_new_biography">
                                <button id="login_button" @click="set_new_biography">
                                    <img id="login_arrow" src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                                </button>
                            </div>
                            <button class="skip_step" @click="skip_registration_step">I don't want to set a biography</button>
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