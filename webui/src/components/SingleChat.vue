<script>
export default {
    props: [
        "chatId",
        "chatName",
        "lastMessageBody",
        "lastMessageSender",
        "lastMessageDate"
    ],

    data: function () {
        return {
            "to_render": true,
            "lastMessageFormattedDate": null
        }
    },

    async mounted() {
        if (!this.lastMessageBody.length || !this.lastMessageSender === null) {
            this.to_render = false;
        }

        let messageTimestamp = new Date(Date.parse(this.lastMessageDate));
        this.lastMessageFormattedDate = `${messageTimestamp.getHours().toString().padStart(2, '0')}:${messageTimestamp.getMinutes().toString().padStart(2, '0')}`;
    }
}
</script>

<template>
    <div class="single_chat" role="button">
        <div class="chat_photo"></div>
        <div class="chat_details">
            <p class="chat_name">{{ chatName }}</p>
            <div class="chat_last_msg_date">
                <p class="chat_last_message" v-if="this.to_render"><b style="font-weight: 650;">{{ lastMessageSender }}</b>: {{ lastMessageBody }}</p>
                <p class="chat_date" v-if="this.to_render">({{ lastMessageFormattedDate }})</p>
            </div>
        </div>
    </div>
</template>