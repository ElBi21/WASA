<script>
import MessageCheck from "@/components/MessageCheck.vue";

export default {
    components: {MessageCheck},

    data: function() {
        return {
            messageFormattedDate: null,
            senderPFP: null
        }
    },

    async mounted() {
        let messageTimestamp = new Date(Date.parse(this.messageObj.timestamp));
        this.messageFormattedDate = `${messageTimestamp.getHours().toString().padStart(2, '0')}:${messageTimestamp.getMinutes().toString().padStart(2, '0')}`;

        this.senderPFP = `data:image/jpeg;base64,` + this.messageObj.sender.profile_pic;
    },

    props: [ "userLogged", "messageObj", "isChatPrivate", "chatUsers" ]
}
</script>

<template>
<div :class="['message_container', { active: messageObj.sender.user_id === userLogged }]">
    <div class="message_side_container message_side_logged_user" v-if="messageObj.sender.user_id === userLogged">
        <div class="message_side_button msg_btn_top" role="button">
            <img src="../assets/icons/trash-can-solid-full.svg" class="msg_side_icon" alt="Edit the message">
        </div>
        <div class="message_side_button" role="button">
            <img src="../assets/icons/arrows-turn-right-solid-full.svg" class="msg_side_icon" alt="Forward the message">
        </div>
        <div class="message_side_button msg_btn_bottom" role="button">
            <img src="../assets/icons/heart-circle-plus-solid-full.svg" class="msg_side_icon" alt="React to the message">
        </div>
    </div>
    <div :class="['message_shape', { active: messageObj.sender.user_id === userLogged }]">
        <div class="message_sender_info" v-if="!isChatPrivate">
            <img class="message_sender_photo" v-if="messageObj.sender.user_id !== userLogged"
                 :src="this.senderPFP" alt="User PFP">
            <p class="message_sender_display_name" v-if="messageObj.sender.user_id !== userLogged">{{ this.messageObj.sender.display_name }}</p>
        </div>
        <p class="message_content">{{ this.messageObj.content }}</p>
        <div class="message_metadata">
            <p class="message_time">{{ this.messageFormattedDate }}</p>
            <MessageCheck :recvList="messageObj.received" :seenList="messageObj.seen" :userList="chatUsers"></MessageCheck>
        </div>
    </div>
    <div class="message_side_container" v-if="messageObj.sender.user_id !== userLogged">
        <div class="message_side_button msg_btn_top" role="button">
            <img src="../assets/icons/arrows-turn-right-solid-full.svg" class="msg_side_icon" alt="Forward the message">
        </div>
        <div class="message_side_button msg_btn_bottom" role="button">
            <img src="../assets/icons/heart-circle-plus-solid-full.svg" class="msg_side_icon" alt="React to the message">
        </div>
    </div>
</div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");

.message_side_container {
    height: 100%;
    width: 40px;

    margin: 0 8px 0 8px;

    display: flex;
    flex-direction: column;
    justify-content: start;
    align-items: start;
    gap: 0;
}

.message_side_logged_user {
    align-items: end;
}

.message_side_button {
    width: 20px;
    height: 20px;
    background-color: var(--glass-background);
    transition: background-color 0.2s ease-in;

    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;

    /* box-shadow: 0 0 4px rgba(0, 0, 0, 0.2); */

    /* border-radius: 8px; */
}

.message_side_button:hover {
    background-color: var(--glass-background-hover);
    transition: background-color 0.2s ease-in;
}

/* From other users */

.msg_btn_top {
    border-radius: 8px 8px 0 0;
}

.msg_btn_bottom {
    border-radius: 0 0 8px 8px;
}

.msg_side_icon {
    width: 12px;
    height: 12px;
    margin: 0;
}

.message_container:hover .message_side_container {
    display: flex;
}
</style>