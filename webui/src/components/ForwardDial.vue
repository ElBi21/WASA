<script>
import {API_get_all_users, API_get_conversations} from "../services/user-ops";
import {API_add_to_group, API_create_conversation, API_get_conversation} from "../services/chat-ops";
import {img_to_base64, retrieveFromStorage} from "../services/utils";
import UserButton from "./UserButton.vue";
import defaultGroupPicture from "../assets/defaults/default_group.png";
import ChatButton from "@/components/ChatButton.vue";

export default {
    components: {ChatButton, UserButton},

    data: function() {
        return {
            userID: null,
            allUsers: [],
            availableUsers: [],

            allChats: [],
            chatsFlags: [],

            usersToAdd: [],
            usersFlags: [],

            safeToCreate: false
        }
    },

    emits: [ "closeNewChatDial" ],

    methods: {
        async toggle_chat_in_conversation(userToCheck) {
            let isUserAddedAlready = this.allChats.findIndex(user => user.user_id === userToCheck.user_id);
            let flagIndex = this.chatsFlags.findIndex(user => user.user_id === userToCheck.user_id);

            // If the user is already added
            if (isUserAddedAlready !== -1) {
                this.usersToAdd.splice(isUserAddedAlready, 1);
                this.usersFlags[flagIndex] = false;
            } else {
                this.usersToAdd.push(userToCheck);
                this.usersFlags[flagIndex] = true;
            }

            this.check_chat_flags();
        },

        check_chat_flags() {
            this.safeToCreate = this.usersToAdd.length >= 1;
        },

        async add_selected_users() {
            for (let user of this.usersToAdd) {
                await API_add_to_group(this.chatID, user.user_id, this.userID);
            }

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
                this.chatsFlags.push(false);
            }
        }


        /* let chatUsers = new Set(chatObj.Users.map(user => user.user_id));
        this.availableUsers = this.allUsers.filter(user => !chatUsers.has(user.user_id));

        for (let _ of this.availableUsers) {
            this.usersFlags.push(false);
        } */
    },

    props: [ "chatID" ],
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
        <ChatButton v-if="allChats.length > 0"
                    v-for="[index, chat] in allChats.entries()"
                    :chat-object="chat"
                    :is-selected="chatsFlags[index]"
                    @click="toggle_chat_in_conversation(chat, index)"></ChatButton>
        <p v-else>Seems like you don't have other chats open at the moment. Create a chat first</p>
    </div>

    <div v-if="safeToCreate === true" class="forward_msg_button_container">
        <button class="forward_msg_confirm_button" style="width: 225px" @click="add_selected_users">
            Forward message
        </button>
    </div>
</div>
</template>

<style scoped>
@import url("../assets/css/forward_msg_dial.css");
</style>