<script>
import {API_get_conversation} from "../services/chat-ops";
import {retrieveFromStorage} from "../services/utils";

export default {
    data: function () {
        return {
            userData: null,
            to_render: true,

            chatName: '',
            chatPicture: '',
            lastMessageFormattedDate: null
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        await this.buildChatButton();
    },

    methods: {
        async buildChatButton() {
            if (!this.lastMessageBody.length || !this.lastMessageSender === null) {
                this.to_render = false;
            }

            let messageTimestamp = new Date(Date.parse(this.lastMessageDate));
            this.lastMessageFormattedDate =
                `${messageTimestamp.getHours().toString().padStart(2, '0')}:${messageTimestamp.getMinutes().toString().padStart(2, '0')}`;

            let chat = await API_get_conversation(this.chatId, this.userData.user_id);

            if (chat.IsPrivate === true) {
                let otherParticipantIndex = chat.Users.findIndex(user => user.user_id !== this.userData.user_id);
                this.chatName = chat.Users[otherParticipantIndex].display_name;
                this.chatPicture = `data:image/jpeg;base64,` + chat.Users[otherParticipantIndex].profile_pic;
            } else {
                this.chatName = chat.Name;
                this.chatPicture = `data:image/jpeg;base64,` + chat.Photo;
            }
        }
    },

    props: [
        "chatId",
        "lastMessageBody",
        "lastMessageSender",
        "lastMessageDate",
        "isChatSelected"
    ],

    watch: {
        async chatId() {
            await this.buildChatButton();
        }
    }
}
</script>

<template>
    <div :class="['single_chat', { active: isChatSelected }]" role="button">
        <img :src="chatPicture" v-if="chatPicture !== ''" class="chat_photo" alt="Chat picture">
        <div class="chat_details">
            <p class="chat_name">{{ chatName }}</p>
            <div class="chat_last_msg_date">
                <p class="chat_last_message" v-if="this.to_render"><b style="font-weight: 650;">{{ lastMessageSender }}</b>: {{ lastMessageBody }}</p>
                <p class="chat_date" v-if="this.to_render">({{ lastMessageFormattedDate }})</p>
            </div>
        </div>
    </div>
</template>