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
            // currentChat: null,
            chatObj: null,
            userData: null
        }
    },

    methods: {
        async updateChatObject() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
    },

    props: [ "selectedChatId" ],

    watch: {
        async selectedChatId() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
        }
    }
}
</script>

<template>
    <div class="chat_window">
        <ChatTopBar v-if="chatObj" :chatObj="chatObj"></ChatTopBar>
        <div class="scroll_container">
            <ChatMessages v-if="chatObj" :chatObj="chatObj"></ChatMessages>
        </div>
        <TypingBar v-if="chatObj" :chatID="selectedChatId"></TypingBar>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");

.scroll_container {
    height: calc(100% + 108px);
    z-index: 1;

    margin: -124px 0 0 0;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    overflow-y: visible;
    overflow-x: hidden;
}
</style>
