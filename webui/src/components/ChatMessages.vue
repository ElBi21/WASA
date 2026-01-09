<script>
import SingleMessage from "./SingleMessage.vue";
import {API_get_conversation_messages} from "@/services/chat-ops";
import {retrieveFromStorage} from "@/services/utils";

export default {
    components: { SingleMessage },

    data: function () {
        return {
            userData: null,
            messages: null,
            reactions: null,

            refresh_timer_ID: null,
            refresh_timer_interval: 2500,

            stopReloading: false,
        }
    },

    emits: [ "openForwardDial", "startForwardToMessage" ],

    methods: {
        async refreshMessages() {
            this.messages = await API_get_conversation_messages(this.chatObj.ID, this.userData.user_id);
        },

        openForwardDial(messageObj) {
            this.$emit("openForwardDial", messageObj);
        },

        startReplyToMessage(messageObj) {
            this.$emit("startForwardToMessage", messageObj);
        },

        async timerRefresh() {
            this.refresh_timer_ID = setInterval(async () => {
                if (this.stopReloading) {
                    this.stopReloading = !this.stopReloading;
                    clearInterval(this.refresh_timer_ID);
                    this.refresh_timer_interval = null;
                } else {
                    await this.refreshMessages();
                }
            }, this.refresh_timer_interval);
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        await this.refreshMessages();

        await this.timerRefresh();
    },

    props: [ "chatObj", "refreshFlag", "stopRefreshFlag", "refreshUser" ],

    watch: {
        async refreshFlag() {
            await this.refreshMessages();
        },

        stopRefreshFlag() {
            this.stopReloading = true;
        },

        async refreshUser() {
            this.userData = await retrieveFromStorage();
            this.stopReloading = true;
            this.stopReloading = false;

            await this.timerRefresh();
        }
    }
}
</script>

<template>
<div class="chat_messages_container">
    <SingleMessage v-for="message in messages"
       @refreshChat="refreshMessages" @openForwardDial="openForwardDial" @startReplyToMessage="startReplyToMessage"
       :user-logged="userData" :message-obj="message" :reactionsObj="message.reactions"
       :chatUsers="chatObj.Users" :isChatPrivate="chatObj.IsPrivate"></SingleMessage>
</div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");
</style>