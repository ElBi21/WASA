<script>
import {API_get_all_users, API_get_conversations} from "../services/user-ops";
import {API_add_to_group, API_create_conversation, API_get_conversation} from "../services/chat-ops";
import {img_to_base64, retrieveFromStorage} from "../services/utils";
import UserButton from "./UserButton.vue";
import defaultGroupPicture from "../assets/defaults/default_group.png";

export default {
    components: {UserButton},

    data: function() {
        return {
            userID: null,
            allUsers: [],
            availableUsers: [],

            usersToAdd: [],
            usersFlags: [],

            safeToCreate: false
        }
    },

    emits: [ "closeNewChatDial" ],

    methods: {
        async toggle_user_in_conversation(userToCheck) {
            let isUserAddedAlready = this.usersToAdd.findIndex(user => user.user_id === userToCheck.user_id);
            let flagIndex = this.availableUsers.findIndex(user => user.user_id === userToCheck.user_id);

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

        let chatObj = await API_get_conversation(this.chatID, this.userID);
        this.allUsers = await API_get_all_users(this.userID);

        let chatUsers = new Set(chatObj.Users.map(user => user.user_id));
        this.availableUsers = this.allUsers.filter(user => !chatUsers.has(user.user_id));

        for (let _ of this.availableUsers) {
            this.usersFlags.push(false);
        }
    },

    props: [ "chatID" ],
}
</script>

<template>
<div class="new_chat_dial">
    <div class="new_chat_close_div">
        <div class="new_chat_close_button" role="button" @click="close_dial">
            <img src="../assets/icons/xmark-solid-full.svg" alt="close" class="new_chat_close_img">
        </div>
    </div>
    <h2 class="new_chat_title">Add a user</h2>

    <div class="new_chat_users_list">
        <div class="extend_full" v-if="availableUsers.length > 0">
            <UserButton v-for="[index, user] in availableUsers.entries()"
                        :key="user.user_id"
                        :user-object="user"
                        :kind-of-chat="selectedOption"
                        :is-selected="usersFlags[index]"
                        @click="toggle_user_in_conversation(user, index)"></UserButton>
        </div>
        <p v-else>Seems like no one else uses WASAText yet. Make someone join in order to add them in this group chat</p>
    </div>

    <div v-if="safeToCreate === true" class="new_chat_button_container">
        <button class="new_chat_confirm_button" style="width: 225px" @click="add_selected_users">
            Add selected user(s)
        </button>
    </div>
</div>
</template>

<style scoped>
@import url("../assets/css/new_chat_dial.css");
</style>