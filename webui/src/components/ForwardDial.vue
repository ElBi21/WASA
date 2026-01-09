<script>
import {API_get_all_users, API_get_conversations} from "../services/user-ops";
import {API_add_to_group, API_create_conversation, API_get_conversation} from "../services/chat-ops";
import {img_to_base64, retrieveFromStorage} from "../services/utils";
import UserButton from "./UserButton.vue";
import defaultGroupPicture from "../assets/defaults/default_group.png";
import ChatButton from "@/components/ChatButton.vue";
import {API_forward_message} from "@/services/message-ops";

export default {
    components: {ChatButton},

    data: function() {
        return {
            userID: null,
            allChats: [],

            selectedChatID: null,

            safeToCreate: false
        }
    },

    emits: [ "closeNewChatDial" ],

    methods: {
        async toggle_chat_in_conversation(selectedChat) {
            this.selectedChatID = selectedChat.ID === this.selectedChatID ? null : selectedChat.ID;
            this.check_chat_flags();
        },

        check_chat_flags() {
            this.safeToCreate = this.selectedChatID !== null;
        },

        async forward_message() {
            await API_forward_message(this.messageObj, this.selectedChatID, this.userID);
            this.$emit("closeNewChatDial");
        },

        async close_dial() {
            this.$emit("closeNewChatDial");
        }
    },

    async mounted() {
        let userData = await retrieveFromStorage();
        this.userID = userData.user_id;

        let userChats = await API_get_conversations(this.userID);

        for (let chat of userChats) {
            if (chat.ID !== this.chatID) {
                this.allChats.push(chat);
            }
        }
    },

    props: [ "chatID", "messageObj" ],
}
</script>

<template>
<div class="forward_msg_dial">
    <div class="forward_msg_close_div">
        <div class="forward_msg_close_button" role="button" @click="close_dial">
            <img src="../assets/icons/xmark-solid-full.svg" alt="close" class="forward_msg_close_img">
        </div>
    </div>
    <h2 class="forward_msg_title">Forward message</h2>

    <div class="forward_msg_users_list">
        <div class="chats_container" v-if="allChats.length > 0">
            <ChatButton v-for="chat in allChats"
                        :chat-object="chat"
                        :is-selected="chat.ID === this.selectedChatID"
                        @click="toggle_chat_in_conversation(chat)"></ChatButton>
        </div>
        <p class="forward_no_chats_text" v-else>Seems like you don't have other chats open at the moment. Create a chat first</p>
    </div>

    <div v-if="safeToCreate === true" class="forward_msg_button_container">
        <button class="forward_msg_confirm_button" @click="forward_message">
            Forward message
        </button>
    </div>
</div>
</template>

<style scoped>
@import url("../assets/css/forward_msg_dial.css");
</style>