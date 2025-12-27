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
            replyMessage: null,

            forceRefreshFlag: false,
            stopTimersFlag: 0
        }
    },

    emits: [ "openAddUserDial", "openForwardDial" ],

    methods: {
        async updateChatObject() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
        },

        refreshChatMessages() {
            this.forceRefreshFlag = !this.forceRefreshFlag;
        },

        openAddUserDial() {
            this.$emit("openAddUserDial");
        },

        openForwardDial(messageObj) {
            this.$emit("openForwardDial", messageObj);
        },

        startReplyToMessage(messageObj) {
            this.replyMessage = messageObj;
        },

        stopReply() {
            this.replyMessage = null;
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
    },

    props: [ "selectedChatId", "logOutStopTimer", "refreshChat" ],

    watch: {
        async selectedChatId() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
            this.refreshChatMessages();
        },

        async refreshChat() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
            this.refreshChatMessages();
        },

        async logOutStopTimer() {
            this.stopTimersFlag = 1;
        }
    }
}
</script>

<template>
    <div class="chat_window">
        <ChatTopBar v-if="chatObj" :chatObj="chatObj" @openAddUserDial="openAddUserDial"></ChatTopBar>
        <div class="scroll_container">
            <ChatMessages v-if="chatObj" :chatObj="chatObj" @openForwardDial="openForwardDial"
                          @startForwardToMessage="startReplyToMessage"
                          :refreshFlag="forceRefreshFlag" :stopRefreshFlag="stopTimersFlag"></ChatMessages>
        </div>
        <TypingBar v-if="chatObj" :chatID="selectedChatId" :replyMessage="replyMessage"
                   @refresh_chat_view="refreshChatMessages" @stop_reply="stopReply"></TypingBar>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>
