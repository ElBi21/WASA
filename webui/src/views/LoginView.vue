<script>
import textJson from "../assets/texts/login.json";
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
            // Flag changing depending on the length of the username inserted
            if (!this.login_valid) {
                // TODO: Implement error message
                return
            }

            // Perform login, wait for promise
            await API_login(this.username).then((value) => {
                Object.assign(this.userData, value.userData);
                this.should_register = value.shouldRegister;

                // Save in local storage
                sessionStorage.setItem("user", JSON.stringify(value.userData));
            })

            // If the user is new, perform registration
            if (this.should_register === true) {
                this.current_register_step += 1;
                await register_ops.start_registration();
            } else {
                this.$router.push({ path: "/home", replace: true });
            }
        },

        // Wrapper for setting the new display name
        async set_new_display_name() {
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
            await register_ops.color_new_progress_bar(this.current_register_step);
            this.current_register_step += 1;

            await API_set_new_pfp(this.userData.user_id, this.userData.profile_pic).then((_) => {
                // Save to local storage and go to home
                sessionStorage.setItem("user", JSON.stringify(this.userData));
            });

            this.$router.push({ path: "/home", replace: true });
        },

        // Event to check if the image is convertible in base64
        async upload_pfp(event) {
            if (!event.target.files.length) return;

            // Save pfp and convert it to base64
            this.newPfp = event.target.files[0];

            let reader = new FileReader();
            let base64Pfp = null;
            reader.onloadend = async () => {
                base64Pfp = reader.result
                    .replace("data:", "")
                    .replace(/^.+,/, "");

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

            // If it's skipping the PFP
            if (this.current_register_step === 3) {
                this.userData.profile_pic = "iVBORw0KGgoAAAANSUhEUgAAAZAAAAGQCAMAAAC3Ycb+AAAAAXNSR0IB2cksfwAAAAlwSFlzAAALEwAACxMBAJqcGAAAAadQTFRFy9PZzNTa09rf2+Dk4ubq6ezv7/L08vT18/X39ff49/j58/X27vHz6Ovu4OXp2d/j0tne0djd3+Tn/Pz8/////v7/+/z87O/x3OHlz9bc4eXp/f397/Hz3OLmz9fc1dvg6u7w+Pn65+rt193i/v7+7vHy0dje8PL05+vuzdXb5urt/Pz9+fr73+To9vj58/T2/P39ztXb4+frzdTa8fP11Nvg9/n58vT2ztbcztbb9fb35+vt7fDy3uPn4ebp1NrfzdXa5Ojr4OXo/f3++vv82uDkzNTZ1tzh+/v84+jr2N7j4+fq4ufq2N7i3eLm2d/k+vv79fb46Ozu/f7+9vf45ens9Pb35uns3OHm6u3w7/Lz6+7x6u3v2+Hl9PX30Nfd1t3h2t/k197i0Nfc+fr6193h7e/yy9Pay9TZytLYy9LY4OPoytLZytLX7vDy5urs1dzg8fT1y9PY2d7iytPZytPYzNPa1Nvh1t3iy9LZ9/n66e3v09ne3+Xoy9HY6+7w3eHm2d7j8PP04+bq4ebq2+Dl4OXn1dzh3+Xp5Ojsy9Ta/v//zNPZve2nBQAAE/hJREFUeJztnXt0FcUZwHcvJIFgkhtIIUAkgSgQBDQCIsrbCopiwT481dNaT2uRqq3WV62n9VEf1dr2tNq3tae1DxVrPcXHQY5tCvWBxaKoBWlVVEREHvJIwpvm5u7MfHvv3rt3d2dnvlm/3z/Z3Gwmm/u7u9/Mt9/M2haBClv3ARBuSAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgySAgy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";
                this.photo_valid = true;
                await this.set_new_pfp();
            }
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
                                   @keyup.enter="login" @input="check_enter_color">
                            <button class="login_button_general login_button-unavailable" @click="login">
                                <img class="login_arrow login_arrow-unavailable"
                                     src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                            </button>
                        </div>
                        <div id="new_display_form">
                            <div class="form_zone">
                                <input v-model="newDisplayName" class="login_register_input" @input="check_enter_color"
                                       placeholder="Choose a display name" @keyup.enter="set_new_display_name">
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
                                       placeholder="Choose a biography" @keyup.enter="set_new_biography" @input="check_enter_color">
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