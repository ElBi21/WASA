<script>
import {img_to_base64, retrieveFromStorage} from "../services/utils";
import {API_send_message} from "../services/message-ops";

export default {
    data: function() {
        return {
            userID: '',
            messageBuffer: '',
            messagePhoto: '',

            isReplying: false,
            replyingTo: '',

            canSend: false
        }
    },

    emits: [ "refresh_chat_view", "stop_reply" ],

    methods: {
        async send_message() {
            this.messageBuffer = this.messageBuffer.trimStart();
            this.messageBuffer = this.messageBuffer.trimEnd();

            await API_send_message(this.userID, this.chatID, this.messageBuffer,
                this.messagePhoto, this.isReplying ? this.replyMessage.message_id : 0);

            this.messageBuffer = '';
            this.messagePhoto = '';
            this.stopReply();

            this.$emit("refresh_chat_view");
        },

        stopReply() {
            this.$emit("stop_reply");
        },

        async load_message_picture(event) {
            let picture = event.target.files[0];
            this.messagePhoto = await img_to_base64(picture);

            // let photoButton = document.getElementById("loaded_pic");
            // photoButton.style.backgroundImage = `data:image/jpeg;base64,` + this.messagePhoto;
        },

        removeImage() {
            this.messagePhoto = '';
        }
    },

    async mounted() {
        let userData = await retrieveFromStorage();
        this.userID = userData.user_id;
    },

    props: [ "chatID", "replyMessage", "refreshUser" ],

    watch: {
        messageBuffer() {
            let controlBuffer = this.messageBuffer.trimStart();
            controlBuffer = controlBuffer.trimEnd();

            this.canSend = controlBuffer.length >= 1;
        },

        chatID() {
            this.messageBuffer = '';
        },

        replyMessage() {
            this.replyingTo = this.replyMessage !== null ? this.replyMessage.content : '';
            this.isReplying = this.replyMessage !== null;
        },

        async refreshUser() {
            let userData = await retrieveFromStorage();
            this.userID = userData.user_id;
        }
    }
}
</script>

<template>
    <div class="typing_bar_container">
        <div class="reply_container" v-if="replyMessage !== null">
            <p class="reply_text"><b>Replying to:</b> {{ replyingTo }}</p>
            <div role="button" class="reply_close" @click="stopReply">
                <img src="../assets/icons/xmark-solid-full.svg" class="reply_close_img" alt="Close reply">
            </div>
        </div>
        <input class="type_bar" placeholder="What do you feel like sharing?" v-model="messageBuffer"
            v-on:keyup.enter="send_message">
        <div v-if="messagePhoto !== ''" class="bar_button" id="loaded_pic" role="button" @click="removeImage">
            <img class="photo_remove" src="../assets/icons/xmark-solid-full.svg" alt="Remove the picture">
            <img class="loaded_pic_img" :src="`data:image/jpeg;base64,` + messagePhoto" alt="Uploaded picture to send">
        </div>
        <label class="bar_button" id="photo_button" role="button">
            <img class="photo_button_icon" src="../assets/icons/images-solid-full.svg" alt="Send photo">
            <input type="file" style="display: none;"
                   accept="image/*" @change="load_message_picture">
        </label>
        <div class="bar_button" :id="canSend === true ? 'send_button_available' : 'send_button'"
             role="button" @click="send_message">
            <img class="send_button_icon" src="../assets/icons/send-solid-full.svg" alt="Send message">
        </div>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>