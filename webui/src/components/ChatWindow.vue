<script setup>
import TypingBar from "./TypingBar.vue";
import ChatTopBar from "./ChatTopBar.vue";
</script>

<script>
export default {
    data: function() {
        return {
            currentChat: null
        }
    },

    props: [ "selectedChatId" ],

    watch: {
        selectedChatID(chat) {
            this.currentChat = chat;
        }
    }


}
</script>

<template>
    <div class="chat_window">
        <ChatTopBar :chatId="selectedChatId"></ChatTopBar>
        <TypingBar></TypingBar>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>
