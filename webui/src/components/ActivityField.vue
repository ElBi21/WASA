<script setup>
import NoChatOpenPage from "@/components/NoChatOpenPage.vue";
import ChatWindow from "@/components/ChatWindow.vue";
import NewChat from "@/components/NewChat.vue";
import AddUsersDial from "@/components/AddUsersDial.vue";
import ForwardDial from "@/components/ForwardDial.vue";
import EditGroupDial from "@/components/EditGroupDial.vue";
import EditUserDial from "@/components/EditUserDial.vue";
</script>

<script>
export default {
    data: function() {
        return {
            selectedChatID: null,
            messageToForward: null,
            openChatObj: null,

            showNoChatPage: false,
            openNewChatDialFlag: false,
            addUserDialFlag: false,
            forwardMessageDial: false,
            openEditGroupDialFlag: false,
            openEditUserDialFlag: false,
            stopTimersFlag: 0,
            refreshChatCounter: 0,
            refreshUserCounter: 0
        }
    },

    emits: [ "closeChatDialExternal", "refreshChatList", "closeEditUserExternal", "refreshUser" ],

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

        closeEditGroupDial() {
            this.openEditGroupDialFlag = false;
            this.refreshChatCounter += 1;
            this.$emit("refreshChatList");
        },

        closeEditUserDial() {
            this.openEditUserDialFlag = false;
            this.$emit("closeEditUserExternal");
        },

        openNewChatDial() {
            this.openNewChatDialFlag = true;
        },

        openAddUserDial() {
            this.addUserDialFlag = true;
        },

        openEditGroupDial(chatObj) {
            this.openChatObj = chatObj;
            this.openEditGroupDialFlag = true;
        },

        async openForwardDial(messageObj) {
            this.forwardMessageDial = true;
            this.messageToForward = messageObj;
        },

        openEditUserDial() {
            this.openEditUserDialFlag = true;
        },

        newUserData(newData) {
            this.refreshUserCounter += 1;
            this.$emit("refreshUser", newData);
        },

        leaveGroup() {
            this.selectedChatID = null;
        }
    },

    props: [ "selectedChatIdProp", "openChatDialExternal", "logOutStopTimer", "openEditUserDialExt" ],

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
        <div v-if="openNewChatDialFlag || openChatDialExternal" class="dial_container">
            <NewChat @close-new-chat-dial="closeNewChatDial"></NewChat>
        </div>

        <div v-if="addUserDialFlag" class="dial_container">
            <AddUsersDial @close-new-chat-dial="closeAddUserDial" :chatID="selectedChatID"></AddUsersDial>
        </div>

        <div v-if="forwardMessageDial" class="dial_container">
            <ForwardDial @close-new-chat-dial="closeForwardDial"
                         :chatID="selectedChatID" :messageObj="messageToForward"></ForwardDial>
        </div>

        <div v-if="openEditGroupDialFlag" class="dial_container">
            <EditGroupDial @closeEditGroupDial="closeEditGroupDial" :chatObj="openChatObj"></EditGroupDial>
        </div>

        <div v-if="openEditUserDialFlag || openEditUserDialExt" class="dial_container">
            <EditUserDial @closeEditUserDial="closeEditUserDial" @newUserValues="newUserData"></EditUserDial>
        </div>

        <NoChatOpenPage v-if="selectedChatID === null || showNoChatPage === true" :refreshUserData="refreshUserCounter"
            @open-new-chat-dial="openNewChatDial" @openEditUserDial="openEditUserDial"></NoChatOpenPage>
        <ChatWindow v-else :selectedChatId="selectedChatID" :logOutStopTimer="stopTimersFlag"
                    :refreshChat="refreshChatCounter" @openEditGroupDial="openEditGroupDial"
                    @openAddUserDial="openAddUserDial" @openForwardDial="openForwardDial"
                    @groupLeft="leaveGroup"></ChatWindow>
    </div>
</template>

<style scoped>
@import url("../assets/css/activity.css");
</style>