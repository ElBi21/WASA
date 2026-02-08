<script>
import {img_to_base64, retrieveFromStorage} from "@/services/utils";
import {
    API_get_all_users,
    API_set_new_biography,
    API_set_new_display_name,
    API_set_new_pfp,
    API_set_new_username
} from "@/services/user-ops";

export default {
    data: function() {
        return {
            previousUserID: "",
            previousDisplayName: "",

            userID: "",
            userDisplayName: "",
            userBio: "",
            userPFP: "",

            usernameConflict: false,
            displayNameConflict: false,
        }
    },

    emits: [ "closeEditUserDial", "newUserValues" ],

    methods: {
        async upload_user_picture(event) {
            let picture = event.target.files[0];
            this.userPFP = await img_to_base64(picture);
        },

        async update_profile() {
            let allUsers = await API_get_all_users(this.previousUserID);
            let usersWithSameUsername = allUsers.filter(user => user.user_id === this.userID);
            let usersWithSameDSName = allUsers.filter(user => user.display_name === this.userDisplayName);

            if (usersWithSameUsername.length !== 0 && usersWithSameUsername[0].user_id !== this.previousUserID) {
                this.usernameConflict = true;
                return;
            } else {
                this.usernameConflict = false;
            }

            if (usersWithSameDSName.length !== 0 && usersWithSameDSName[0].user_id !== this.previousUserID) {
                this.displayNameConflict = true;
                return;
            } else {
                this.displayNameConflict = false;
            }

            await API_set_new_username(this.previousUserID, this.userID);
            await API_set_new_display_name(this.userID, this.userDisplayName);
            await API_set_new_biography(this.userID, this.userBio);
            await API_set_new_pfp(this.userID, this.userPFP);

            let userData = {
                user_id: this.userID,
                display_name: this.userDisplayName,
                biography: this.userBio,
                profile_pic: this.userPFP
            };

            sessionStorage.setItem("user", JSON.stringify(userData));

            this.$emit("closeEditUserDial");
            this.$emit("newUserValues", userData);
        },

        async close_dial() {
            this.$emit("closeEditUserDial");
        }
    },

    async mounted() {
        let userData = await retrieveFromStorage();
        this.userID = userData.user_id;
        this.previousUserID = userData.user_id;

        this.userDisplayName = userData.display_name;
        this.previousDisplayName = userData.display_name;

        this.userBio = userData.biography;
        this.userPFP = userData.profile_pic;
    },

    name: "EditUserDial",
}
</script>

<template>
    <div class="edit_group_dial">
        <div class="edit_group_close_div">
            <div class="edit_group_close_button" role="button" @click="close_dial">
                <img src="../assets/icons/xmark-solid-full.svg" alt="close" class="edit_group_close_img">
            </div>
        </div>
        <h2 class="edit_group_title">Edit your profile</h2>
        <div class="edit_group_settings">
            <div class="edit_group_settings_text">
                <div class="edit_username_container">
                    <p class="edit_group_text">Your Username</p>
                    <p v-if="usernameConflict" class="edit_error_text"
                        >This username has already been taken. Please select another one</p>
                </div>

                <input
                    v-model="userID"
                    type="text"
                    placeholder="What's your new username going to be?"
                    class="edit_group_input"
                />

                <div class="edit_username_container">
                    <p class="edit_group_text">Your Display Name</p>
                    <p v-if="displayNameConflict" class="edit_error_text"
                    >This display name has already been taken. Please select another one</p>
                </div>
                <input
                    v-model="userDisplayName"
                    type="text"
                    placeholder="What's your new display name?"
                    class="edit_group_input"
                />

                <p class="edit_group_text">Biography (optional)</p>
                <textarea
                    v-model="userBio"
                    placeholder="Share something with the others"
                    class="edit_group_textarea"
                    rows="2"
                ></textarea>
            </div>

            <div class="edit_group_settings_picture">
                <label class="edit_group_picture_label"
                       :style="{ backgroundImage: `url(data:image/jpeg;base64,${userPFP})` }">
                    Click to upload a group picture
                    <input class="edit_group_picture_button" type="file"
                           accept="image/*" @change="upload_user_picture">
                </label>
            </div>
        </div>

        <div v-if="userID.length >= 1 && userDisplayName.length >= 1" class="edit_group_button_container">
            <button class="edit_group_confirm_button" @click="update_profile">
                Update your profile
            </button>
        </div>
    </div>
</template>

<style scoped>
@import url("../assets/css/edit_dials.css");
</style>