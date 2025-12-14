<script>
import SingleMessage from "./SingleMessage.vue";
import {API_get_conversation_messages} from "../services/chat-ops";
import {retrieveFromStorage} from "../services/utils";
import {API_get_conversations} from "@/services/user-ops";

export default {
    components: { SingleMessage },

    data: function () {
        return {
            userData: null,
            messages: null,

            refresh_timer_ID: null,
            refresh_timer_interval: 2500,

            stopReloading: false
        }
    },

    methods: {
        async refreshMessages() {
            this.messages = await API_get_conversation_messages(this.chatObj.ID, this.userData.user_id);
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        await this.refreshMessages();

        this.refresh_timer_ID = setInterval(async () => {
            if (this.stopReloading) {
                this.stopReloading = !this.stopReloading;
                clearInterval(this.refresh_timer_ID);
                this.refresh_timer_interval = null;
            } else {
                await this.refreshMessages();
            }
        }, this.refresh_timer_interval);
    },

    props: [ "chatObj", "refreshFlag", "stopRefreshFlag" ],

    watch: {
        async refreshFlag() {
            await this.refreshMessages();
        },

        stopRefreshFlag() {
            this.stopReloading = true;
        }
    }
}
</script>

<template>
<div class="chat_messages_container">
    <SingleMessage v-for="message in messages" @refreshChat="refreshMessages"
       :user-logged="userData.user_id" :message-obj="message"
       :chatUsers="chatObj.Users" :isChatPrivate="chatObj.IsPrivate"></SingleMessage>
</div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");
</style>