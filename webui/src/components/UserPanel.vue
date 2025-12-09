<script>
import {clearSessionStorage, retrieveFromStorage} from "../services/utils.js";

export default {
    data: function () {
        return {
            imageSource: "",
            userData: null
        }
    },

    emits: [ "logOutClicked" ],

    methods: {
        async goToLogin() {
            await clearSessionStorage();
            this.$emit("logOutClicked");
            this.$router.push({path: "/", replace: true});
        },

        async prepareUserData() {
            let format = "jpeg";

            if (!this.userData.profile_pic.startsWith('data:image/')) {
                this.userData.profile_pic = `data:image/${format};base64,` + this.userData.profile_pic;
            }

            this.imageSource = this.userData.profile_pic;
        }
    },

    async mounted () {
        this.userData = await retrieveFromStorage();
        await this.prepareUserData();
    }
}
</script>

<template>
    <div class="user_panel">
        <img :src="imageSource" alt="User PFP" class="user_pfp">
        <div class="user_data" v-if="this.userData !== null">
            <p class="user_info"><b>{{ this.userData.display_name }}</b>&nbsp(@{{ this.userData.user_id }})</p>
            <p class="user_bio">{{ this.userData.biography }}</p>
        </div>
        <button class="home_button" id="logout_button" @click="goToLogin">
            <img class="home_button_icon" id="logout_button_icon"
                 src="../assets/icons/door-open-solid-full.svg" alt="Settings">
        </button>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_list.css");
</style>