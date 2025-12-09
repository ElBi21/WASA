<script>
import {retrieveFromStorage} from "../services/utils";
import {API_send_message} from "../services/message-ops";

export default {
    data: function() {
        return {
            userID: '',
            messageBuffer: '',
            messagePhoto: '',
            isReplying: false,

            canSend: false
        }
    },

    emits: [ "refresh_chat_view" ],

    methods: {
        async send_message() {
            this.messageBuffer = this.messageBuffer.trimStart();
            this.messageBuffer = this.messageBuffer.trimEnd();

            await API_send_message(this.userID, this.chatID, this.messageBuffer, this.messagePhoto, this.isReplying);

            this.messageBuffer = '';

            this.$emit("refresh_chat_view");
        }
    },

    async mounted() {
        let userData = await retrieveFromStorage();
        this.userID = userData.user_id;
    },

    props: [ "chatID" ],

    watch: {
        messageBuffer() {
            let controlBuffer = this.messageBuffer.trimStart();
            controlBuffer = controlBuffer.trimEnd();

            this.canSend = controlBuffer.length >= 1;
        },

        chatID() {
            this.messageBuffer = '';
        }
    }
}
</script>

<template>
    <div class="typing_bar_container">
        <input class="type_bar" placeholder="What do you feel like sharing?" v-model="messageBuffer"
            v-on:keyup.enter="send_message">
        <div class="bar_button" id="photo_button" role="button">
            <img class="photo_button_icon" src="../assets/icons/images-solid-full.svg" alt="Send photo">
        </div>
        <div class="bar_button" :id="canSend === true ? 'send_button_available' : 'send_button'"
             role="button" @click="send_message">
            <img class="send_button_icon" src="../assets/icons/send-solid-full.svg" alt="Send message">
        </div>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>