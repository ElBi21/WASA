<script>
import {API_get_all_users, API_get_conversations} from "@/services/user-ops";
import {API_forward_message} from "@/services/message-ops";
import {retrieveFromStorage} from "@/services/utils";
import ChatButton from "@/components/ChatButton.vue";
import UserButton from "@/components/UserButton.vue";
import {API_create_conversation} from "@/services/chat-ops";

export default {
    components: {UserButton, ChatButton},

    data: function() {
        return {
            userID: null,
            allChats: [],
            possibleNewPrivateChats: [],

            selectedChatID: null,
            selectedUserID: null,

            safeToCreate: false
        }
    },

    emits: [ "closeNewChatDial" ],

    methods: {
        toggle_chat_in_conversation(selectedChat) {
            this.selectedChatID = selectedChat.ID === this.selectedChatID ? null : selectedChat.ID;
            this.check_chat_flags();
        },

        other_user_selected(selectedUser) {
            this.selectedUserID = selectedUser.user_id === this.selectedUserID ? null : selectedUser.user_id;
            this.check_chat_flags();
        },

        check_chat_flags() {
            this.safeToCreate = !!(this.selectedChatID !== null ^ this.selectedUserID !== null);
        },

        async forward_message() {
            if (this.selectedChatID === null && this.selectedUserID !== null) {
                let newChatObj = await API_create_conversation(true, [this.userID, this.selectedUserID], '', '', '', this.userID);
                this.selectedChatID = newChatObj.ID;
                this.selectedUserID = null;
            }

            if (this.selectedChatID !== null && this.selectedUserID === null) {
                await API_forward_message(this.messageObj, this.selectedChatID, this.userID);
                this.$emit("closeNewChatDial");
            }
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

        let allUsers = await API_get_all_users(this.userID);

        for (let user of allUsers) {
            let isUserIn = this.allChats.filter(chat =>
                chat.Users.some(usr => usr.user_id === user.user_id) && chat.IsPrivate);

            if (user.user_id !== this.userID && isUserIn.length === 0) {
                this.possibleNewPrivateChats.push(user);
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
            <p class="forward_section_text less_margin_below">Already started chats</p>
            <ChatButton v-for="chat in allChats"
                        :key="chat.ID"
                        :chat-object="chat"
                        :is-selected="chat.ID === this.selectedChatID"
                        @click="toggle_chat_in_conversation(chat)"></ChatButton>
        </div>
        <div class="chats_container" v-if="possibleNewPrivateChats.length > 0">
            <div class="vert_divider"></div>
            <p class="forward_section_text">New users to start</p>
            <UserButton v-for="user in possibleNewPrivateChats"
                        :key="user.user_id"
                        :userObject="user"
                        :isSelected="user.user_id === this.selectedUserID"
                        @click="other_user_selected(user)"></UserButton>
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