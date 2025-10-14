<script setup>
import NoChatOpenPage from "./NoChatOpenPage.vue";
import ChatWindow from "./ChatWindow.vue";
import NewChat from "./NewChat.vue";
</script>

<script>
export default {
    data: function() {
        return {
            selectedChatID: null,
            showNoChatPage: false,

            openNewChatDialFlag: false
        }
    },

    methods: {
        closeNewChatDial() {
            this.openNewChatDialFlag = false;
        },

        async openNewChatDial() {
            console.log(this.openNewChatDialFlag);
            this.openNewChatDialFlag = true;
        }
    },

    props: [ "selectedChatIdProp" ],

    watch: {
        selectedChatIdProp(chatID) {
            this.selectedChatID = chatID;
        }
    }
}
</script>

<template>
    <div class="main_activity_space">
        <div v-if="openNewChatDialFlag === true" class="dial_container">
            <NewChat @close-new-chat-dial="closeNewChatDial"></NewChat>
        </div>

        <NoChatOpenPage v-if="selectedChatID === null || showNoChatPage === true"
            @open-new-chat-dial="openNewChatDial"></NoChatOpenPage>
        <ChatWindow v-else :selected-chat-id="selectedChatID"></ChatWindow>
    </div>
</template>

<style scoped>
@import url("../assets/css/activity.css");
</style>