<script setup>
import TypingBar from "./TypingBar.vue";
import ChatTopBar from "./ChatTopBar.vue";
import ChatMessages from "./ChatMessages.vue";
</script>

<script>
import {API_get_conversation} from "../services/chat-ops";
import {retrieveFromStorage} from "../services/utils";

export default {
    data: function() {
        return {
            chatObj: null,
            userData: null,

            forceRefreshFlag: false,
            stopTimersFlag: 0
        }
    },

    methods: {
        async updateChatObject() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
        },

        refreshChatMessages() {
            console.log(`User ${this.userData.user_id} refreshed the chat`);
            this.forceRefreshFlag = !this.forceRefreshFlag;
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
    },

    props: [ "selectedChatId", "logOutStopTimer" ],

    watch: {
        async selectedChatId() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
            this.refreshChatMessages();
        },

        async logOutStopTimer() {
            this.stopTimersFlag = 1;
            console.log(`It's time to STOP!`);
        }
    }
}
</script>

<template>
    <div class="chat_window">
        <ChatTopBar v-if="chatObj" :chatObj="chatObj"></ChatTopBar>
        <div class="scroll_container">
            <ChatMessages v-if="chatObj" :chatObj="chatObj"
                          :refreshFlag="forceRefreshFlag" :stopRefreshFlag="stopTimersFlag"></ChatMessages>
        </div>
        <TypingBar v-if="chatObj" :chatID="selectedChatId"
                   @refresh_chat_view="refreshChatMessages"></TypingBar>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>
