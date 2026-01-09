<script>
import {retrieveFromStorage} from "@/services/utils";

export default {
    data: function() {
        return {
            user_data: {
                display_name: "",
            }
        }
    },

    emits: [ "openNewChatDial", "openEditUserDial" ],

    methods: {
        openNewChatDialEmit() {
            this.$emit('openNewChatDial');
        },

        openEditUserDial() {
            this.$emit("openEditUserDial");
        }
    },

    async mounted() {
        this.user_data = await retrieveFromStorage();
    },

    props: [ "refreshUserData" ],

    watch: {
        async refreshUserData() {
            let userData = await retrieveFromStorage();
            this.user_data.display_name = userData.display_name;
        }
    }
}
</script>

<template>
    <div class="no_chat_open_main">
        <h1 class="no_chat_title">Welcome {{ this.user_data.display_name }}</h1>
        <p class="no_chat_description"> Hey there, let's do something great today!
            If you're new to WASAText, you should start with some of the suggestions listed here below</p>

        <div class="no_chat_suggestions">
            <div class="no_chat_single_suggestion" id="sugg_new_chat"
                 role="button" @click="openNewChatDialEmit">
                <img class="no_chat_suggestion_icon" src="../assets/icons/chats-solid-full.svg">
                Start a chat</div>
            <div class="no_chat_single_suggestion" id="sugg_prof_settings"
                 role="button" @click="openEditUserDial">
                <img class="no_chat_suggestion_icon" src="../assets/icons/user-solid-full.svg">
                Change your profile</div>
        </div>
    </div>
</template>

<style scoped>
@import url("../assets/css/activity.css");
</style>