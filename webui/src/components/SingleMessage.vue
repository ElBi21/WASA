<script>
export default {
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

    props: [ "userLogged", "messageObj" ]
}
</script>

<template>
<div :class="['message_container', { active: messageObj.sender.user_id === userLogged }]">
    <div :class="['message_shape', { active: messageObj.sender.user_id === userLogged }]">
        <div class="message_sender_info">
            <img class="message_sender_photo" v-if="messageObj.sender.user_id !== userLogged"
                 :src="this.senderPFP" alt="User PFP">
            <p class="message_sender_display_name" v-if="messageObj.sender.user_id !== userLogged">{{ this.messageObj.sender.display_name }}</p>
        </div>
        <p class="message_content">{{ this.messageObj.content }}</p>
        <div class="message_metadata">
            <p class="message_time">{{ this.messageFormattedDate }}</p>
            <div class="message_status"></div>
        </div>
    </div>
</div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");
</style>