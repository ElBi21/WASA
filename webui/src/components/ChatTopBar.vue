<script>
import {API_get_conversation} from "../services/chat-ops";
import {retrieveFromStorage} from "../services/utils";

export default {
    data: function() {
        return {
            userData: null,

            chatTitle: "",
            chatPhoto: null,
            chatUsersString: "",
            chatIsPrivate: false,
        }
    },

    methods: {
        async loadChatData() {
            if (!this.chatObj.IsPrivate) {
                this.chatPhoto = `data:image/jpeg;base64,` + this.chatObj.Photo;

                this.chatTitle = this.chatObj.Name;
                this.chatUsersString = this.chatObj.Users.map(user => user.display_name).join(', ');
            } else {
                let otherUser = this.chatObj.Users[0].user_id === this.userData.user_id ? this.chatObj.Users[1] : this.chatObj.Users[0];
                this.chatPhoto = `data:image/jpeg;base64,` + otherUser.profile_pic;

                this.chatUsersString = "";
                this.chatTitle = otherUser.display_name;
            }
        },
    },

    async mounted() {
        this.userData = await retrieveFromStorage();

        await this.loadChatData();
    },

    props: [ "chatObj" ],

    watch: {
        async chatObj(newValue) {
            await this.loadChatData();
        }
    }
}
</script>

<template>
<div class="chat_top_bar text-wrapper" v-if="chatObj">
    <div class="chat_top_main_info">
        <img :src="chatPhoto" class="chat_top_image" alt="chat pfp">
        <div class="chat_top_info text-wrapper">
            <h1 class="chat_top_title text-wrapper">{{ this.chatTitle }}</h1>
            <p v-if="!this.chatObj.IsPrivate" class="chat_top_paragraph text-wrapper">{{ chatUsersString }}</p>
            <p v-if="!this.chatObj.IsPrivate && this.chatObj.GroupDescription" id="chat_top_desc" class="chat_top_paragraph text-wrapper">{{ chatObj.GroupDescription }}</p>
        </div>
    </div>
    <!--<div class="separator_vertical"></div>-->
</div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>