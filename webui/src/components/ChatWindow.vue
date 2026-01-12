<script>
import TypingBar from "@/components/TypingBar.vue";
import ChatTopBar from "@/components/ChatTopBar.vue";
import ChatMessages from "@/components/ChatMessages.vue";
import {API_get_conversation} from "@/services/chat-ops";
import {retrieveFromStorage} from "@/services/utils";

export default {
    components: {TypingBar, ChatTopBar, ChatMessages},

    data: function() {
        return {
            chatObj: null,
            userData: null,
            replyMessage: null,

            forceRefreshFlag: 0,
            stopTimersFlag: 0,
            stopReloading: false,
            refresh_timer_interval: 1500,
            refresh_timer_ID: null,
        }
    },

    emits: [ "openAddUserDial", "openForwardDial", "openEditGroupDial", "groupLeft" ],

    methods: {
        async updateChatObject() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
        },

        refreshChatMessages() {
            this.forceRefreshFlag++;
        },

        openAddUserDial() {
            this.$emit("openAddUserDial");
        },

        openForwardDial(messageObj) {
            this.$emit("openForwardDial", messageObj);
        },

        openEditGroupDial(chatObj) {
            this.$emit("openEditGroupDial", chatObj);
        },

        startReplyToMessage(messageObj) {
            this.replyMessage = messageObj;
        },

        stopReply() {
            this.replyMessage = null;
        },

        leaveGroup() {
            this.$emit("groupLeft");
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);

        this.refresh_timer_ID = setInterval(async () => {
            if (this.stopReloading && this.chatObj !== null) {
                this.stopReloading = !this.stopReloading;
                clearInterval(this.refresh_timer_ID);
                this.refresh_timer_interval = null;
            } else {
                this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
            }
        }, this.refresh_timer_interval);
    },

    props: [ "selectedChatId", "logOutStopTimer", "refreshChat", "refreshUser" ],

    watch: {
        async selectedChatId() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
            this.refreshChatMessages();
        },

        async refreshChat() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
            this.refreshChatMessages();
        },

        async refreshUser() {
            this.chatObj = await API_get_conversation(this.selectedChatId, this.userData.user_id);
            this.userData = await retrieveFromStorage();
        },

        async logOutStopTimer() {
            this.stopTimersFlag = 1;
        }
    }
}
</script>

<template>
    <div class="chat_window">
        <ChatTopBar v-if="chatObj" :chatObj="chatObj" :refreshUser="refreshUser" @groupLeft="leaveGroup"
                    @openAddUserDial="openAddUserDial" @openEditGroupDial="openEditGroupDial"></ChatTopBar>
        <div class="scroll_container">
            <ChatMessages v-if="chatObj" :chatObj="chatObj" @openForwardDial="openForwardDial"
                          @startForwardToMessage="startReplyToMessage" :refreshUser="refreshUser"
                          :refreshFlag="forceRefreshFlag" :stopRefreshFlag="stopTimersFlag"></ChatMessages>
        </div>
        <TypingBar v-if="chatObj" :chatID="selectedChatId" :replyMessage="replyMessage" :refreshUser="refreshUser"
                   @refresh_chat_view="refreshChatMessages" @stop_reply="stopReply"></TypingBar>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>
