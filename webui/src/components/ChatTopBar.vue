<script>
import {API_get_conversation} from "../services/chat-ops";
import {retrieveFromStorage} from "../services/utils";

export default {
    data: function() {
        return {
            userData: null,
            chatObj: null,
            chatPhoto: null,

            chatUsersString: ""
        }
    },

    methods: {
        async loadChatData(chat_id, user_id) {
            this.chatObj = await API_get_conversation(chat_id, user_id);

            if (!this.chatObj.IsPrivate) {
                this.chatPhoto = `data:image/jpeg;base64,` + this.chatObj.Photo
            }
        },

        async createUsersString() {
            this.chatUsersString = this.chatObj.Users.map(user => user.display_name).join(', ');
        }
    },

    async mounted() {
        let userData = await retrieveFromStorage();
        await this.loadChatData(this.chatId, userData.user_id);
        await this.createUsersString();

        console.log(this.chatUsersString);
    },

    props: [ "chatId" ],

    watch: {
        async chatId(chat_id) {
            let userData = await retrieveFromStorage();
            await this.loadChatData(chat_id, userData.user_id);
            await this.createUsersString();
        }
    }
}
</script>

<template>
<div class="chat_top_bar text-wrapper" v-if="chatObj">
    <img :src="chatPhoto" class="chat_top_image" alt="chat pfp">
    <div class="chat_top_info text-wrapper">
        <div class="chat_top_name_desc">
            <h1 class="chat_top_title text-wrapper">{{ chatObj.Name }}</h1>
            <p class="chat_top_paragraph text-wrapper">{{ chatObj.GroupDescription }}</p>
        </div>
        <p class="chat_top_paragraph text-wrapper">{{ chatUsersString }}</p>
    </div>
    <div class="separator_vertical"></div>
</div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>