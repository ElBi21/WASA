<script>
import MessageCheck from "@/components/MessageCheck.vue";
import {API_delete_message} from "@/services/message-ops";

export default {
    components: {MessageCheck},

    data: function() {
        return {
            messageFormattedDate: null,
            senderPFP: null
        }
    },

    emits: [ "refreshChat" ],

    methods: {
        async deleteMessage() {
            await API_delete_message(this.messageObj.message_id, this.userLogged);
            this.$emit("refreshChat");
        }
    },

    async mounted() {
        let messageTimestamp = new Date(Date.parse(this.messageObj.timestamp));
        this.messageFormattedDate = `${messageTimestamp.getHours().toString().padStart(2, '0')}:${messageTimestamp.getMinutes().toString().padStart(2, '0')}`;

        this.senderPFP = `data:image/jpeg;base64,` + this.messageObj.sender.profile_pic;

        // console.log(this.messageObj, this.userLogged);
    },

    props: [ "userLogged", "messageObj", "isChatPrivate", "chatUsers" ]
}
</script>

<template>
<div :class="['message_container', { active: messageObj.sender.user_id === userLogged }]">
    <div class="message_side_container message_side_logged_user" v-if="messageObj.sender.user_id === userLogged">
        <div class="message_side_button msg_btn_top" role="button" @click="deleteMessage" v-if="!this.messageObj.deleted">
            <img src="../assets/icons/trash-can-solid-full.svg" class="msg_side_icon" alt="Edit the message">
        </div>
        <div class="message_side_button" role="button" v-if="!this.messageObj.deleted">
            <img src="../assets/icons/arrows-turn-right-solid-full.svg" class="msg_side_icon" alt="Forward the message">
        </div>
        <div :class="['message_side_button', this.messageObj.deleted ? 'msg_btn_round' : 'msg_btn_bottom']" role="button">
            <img src="../assets/icons/heart-circle-plus-solid-full.svg" class="msg_side_icon" alt="React to the message">
        </div>
    </div>
    <div :class="['message_shape', { active: messageObj.sender.user_id === userLogged }]">
        <div class="message_sender_info" v-if="!isChatPrivate">
            <img class="message_sender_photo" v-if="messageObj.sender.user_id !== userLogged"
                 :src="this.senderPFP" alt="User PFP">
            <p class="message_sender_display_name" v-if="messageObj.sender.user_id !== userLogged">{{ this.messageObj.sender.display_name }}</p>
        </div>

        <p v-if="!this.messageObj.deleted" class="message_content">{{ this.messageObj.content }}</p>
        <div v-else class="message_deleted">
            <img src="../assets/icons/ban-solid-full.svg" class="message_deleted_icon">
            <p class="message_deleted_text">This message has been deleted</p>
        </div>

        <div class="message_metadata">
            <p class="message_time">{{ this.messageFormattedDate }}</p>
            <MessageCheck :recvList="messageObj.received" :seenList="messageObj.seen" :userList="chatUsers"></MessageCheck>
        </div>
    </div>
    <div class="message_side_container" v-if="messageObj.sender.user_id !== userLogged">
        <div class="message_side_button msg_btn_top" role="button" v-if="!this.messageObj.deleted">
            <img src="../assets/icons/arrows-turn-right-solid-full.svg" class="msg_side_icon" alt="Forward the message">
        </div>
        <div :class="['message_side_button', this.messageObj.deleted ? 'msg_btn_round' : 'msg_btn_bottom']" role="button">
            <img src="../assets/icons/heart-circle-plus-solid-full.svg" class="msg_side_icon" alt="React to the message">
        </div>
    </div>
</div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");
</style>