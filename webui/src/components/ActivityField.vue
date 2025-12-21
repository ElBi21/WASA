<script setup>
import NoChatOpenPage from "./NoChatOpenPage.vue";
import ChatWindow from "./ChatWindow.vue";
import NewChat from "./NewChat.vue";
import AddUsersDial from "@/components/AddUsersDial.vue";
import ForwardDial from "@/components/ForwardDial.vue";
</script>

<script>
import {retrieveFromStorage} from "@/services/utils";

export default {
    data: function() {
        return {
            selectedChatID: null,
            messageToForward: null,

            showNoChatPage: false,
            openNewChatDialFlag: false,
            addUserDialFlag: false,
            forwardMessageDial: false,
            stopTimersFlag: 0,
            refreshChatCounter: 0
        }
    },

    emits: [ "closeChatDialExternal" ],

    methods: {
        closeNewChatDial() {
            this.openNewChatDialFlag = false;
            this.$emit("closeChatDialExternal");
        },

        async closeAddUserDial() {
            this.addUserDialFlag = false;
            this.refreshChatCounter += 1;
        },

        async closeForwardDial() {
            this.forwardMessageDial = false;
            this.messageToForward = null;
        },

        async openNewChatDial() {
            this.openNewChatDialFlag = true;
        },

        async openAddUserDial() {
            this.addUserDialFlag = true;
        },

        async openForwardDial(messageObj) {
            this.forwardMessageDial = true;
            this.messageToForward = messageObj;
        }
    },

    props: [ "selectedChatIdProp", "openChatDialExternal", "logOutStopTimer" ],

    watch: {
        selectedChatIdProp(chatID) {
            this.selectedChatID = chatID;
            this.addUserDialFlag = false;
        },

        logOutStopTimer() {
            this.stopTimersFlag = 1;
        }
    }
}
</script>

<template>
    <div class="main_activity_space">
        <div v-if="openNewChatDialFlag || openChatDialExternal " class="dial_container">
            <NewChat @close-new-chat-dial="closeNewChatDial"></NewChat>
        </div>

        <div v-if="addUserDialFlag" class="dial_container">
            <AddUsersDial @close-new-chat-dial="closeAddUserDial" :chatID="selectedChatID"></AddUsersDial>
        </div>

        <div v-if="forwardMessageDial" class="dial_container">
            <ForwardDial @close-new-chat-dial="closeForwardDial"
                         :chatID="selectedChatID" :messageObj="messageToForward"></ForwardDial>
        </div>

        <NoChatOpenPage v-if="selectedChatID === null || showNoChatPage === true"
            @open-new-chat-dial="openNewChatDial"></NoChatOpenPage>
        <ChatWindow v-else :selectedChatId="selectedChatID" :logOutStopTimer="stopTimersFlag"
                    :refreshChat="refreshChatCounter"
                    @openAddUserDial="openAddUserDial" @openForwardDial="openForwardDial"></ChatWindow>
    </div>
</template>

<style scoped>
@import url("../assets/css/activity.css");
</style>