<script>
import {API_get_conversation} from "../services/chat-ops";
import {retrieveFromStorage} from "../services/utils";

export default {
    props: [
        "chatId",
        "lastMessageBody",
        "lastMessageSender",
        "lastMessageDate"
    ],

    data: function () {
        return {
            to_render: true,

            chatName: '',
            lastMessageFormattedDate: null
        }
    },

    async mounted() {
        let userData = await retrieveFromStorage();

        if (!this.lastMessageBody.length || !this.lastMessageSender === null) {
            this.to_render = false;
        }

        let messageTimestamp = new Date(Date.parse(this.lastMessageDate));
        this.lastMessageFormattedDate = `${messageTimestamp.getHours().toString().padStart(2, '0')}:${messageTimestamp.getMinutes().toString().padStart(2, '0')}`;

        let chat = await API_get_conversation(this.chatId, userData.user_id);

        if (chat.IsPrivate === true) {
            let otherParticipantIndex = chat.Users.findIndex(user => user.user_id !== userData.user_id);
            this.chatName = chat.Users[otherParticipantIndex].display_name;
        } else {
            this.chatName = chat.Name;
        }
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