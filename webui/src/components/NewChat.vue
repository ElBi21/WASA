<script>
import {API_get_all_users, API_get_conversations} from "../services/user-ops";
import {API_create_conversation} from "../services/chat-ops";
import {img_to_base64, retrieveFromStorage} from "../services/utils";
import UserButton from "./UserButton.vue";
import defaultGroupPicture from "../assets/defaults/default_group.png";

export default {
    components: {UserButton},

    data: function() {
        return {
            userID: null,
            allUsers: [],
            privateChatUsers: [],
            selectedOption: 'private',

            newGroupName: "",
            newGroupDescription: "",
            newGroupPhoto: "",

            usersToAdd: [],

            usersGroupFlags: [],
            usersPrivateFlags: [],

            safeToCreate: false,
            readyToDraw: false
        }
    },

    emits: [ "closeNewChatDial" ],

    methods: {
        async upload_group_picture(event) {
            let picture = event.target.files[0];
            this.newGroupPhoto = await img_to_base64(picture);
        },

        async toggle_user_in_conversation(userParam, flagIndex) {
            let isUserAddedAlready = this.usersToAdd.findIndex(user => user === userParam.user_id);

            // If the user is already added
            if (isUserAddedAlready !== -1) {
                this.usersToAdd.splice(isUserAddedAlready, 1);

                let indexInPeopleArray = this.allUsers.findIndex(user => user.user_id === userParam.user_id);
                let indexInPrivateArray = this.privateChatUsers.findIndex(user => user.user_id === userParam.user_id);
                this.selectedOption === 'private' ?
                    this.usersPrivateFlags[indexInPrivateArray] = false :
                    this.usersGroupFlags[indexInPeopleArray] = false;

                this.safeToCreate = await this.check_chat_flags();

                return;
            }

            // If chat is private, check if any user is already selected. If any is, then remove it
            if (this.selectedOption === "private") {
                // There is someone else selected
                if (this.usersToAdd.length > 0) {
                    // Reset all users
                    this.usersToAdd = []
                    for (let i = 0; i < this.usersPrivateFlags.length; i++) {
                        this.usersPrivateFlags[i] = false;
                    }
                }
            }

            this.usersToAdd.push(userParam.user_id);
            this.selectedOption === 'private' ?
                this.usersPrivateFlags[flagIndex] = true :
                this.usersGroupFlags[flagIndex] = true;

            this.safeToCreate = await this.check_chat_flags();
        },

        async check_chat_flags() {
            return (this.selectedOption === 'group' && (this.newGroupName.length >= 1 && this.usersToAdd.length >= 1))
                || (this.selectedOption === 'private' && this.usersToAdd.length === 1);
        },

        async create_new_conversation() {
            let privateFlag = this.selectedOption === 'private';

            if (this.selectedOption === 'private') {
                this.newGroupName = '';
                this.newGroupDescription = '';
                this.newGroupPhoto = '';
            }

            this.usersToAdd.push(this.userID);

            await API_create_conversation(
                privateFlag,
                this.usersToAdd,
                this.newGroupName,
                this.newGroupDescription,
                this.newGroupPhoto,
                this.userID
            );

            this.$emit("closeNewChatDial");
        },

        async close_dial() {
            this.$emit("closeNewChatDial");
        }
    },

    async mounted() {
        let userData = await retrieveFromStorage();
        this.userID = userData.user_id;

        let userConvs = await API_get_conversations(this.userID);
        this.allUsers = await API_get_all_users(this.userID);

        let index = 0;
        let continueIncreasing = true;
        for (let user of this.allUsers) {
            if (user.user_id === this.userID) {
                continueIncreasing = false;
            }

            if (continueIncreasing) {
                index++;
            }

            this.usersGroupFlags.push(false);
        }

        this.allUsers.splice(index, 1);
        this.usersGroupFlags.splice(0, 1);

        let usersToRemove = new Set();
        if (userConvs !== null) {
            for (let chat of userConvs) {
                if (chat.IsPrivate === true) {
                    for (let user of chat.Users) {
                        if (user.user_id !== this.userID) {
                            usersToRemove.add(user.user_id);
                        }
                    }
                }
            }

            this.privateChatUsers = this.allUsers.filter(user => !usersToRemove.has(user.user_id));
        } else {
            this.privateChatUsers = this.allUsers;
        }

        for (let _ of this.privateChatUsers) {
            this.usersPrivateFlags.push(false);
        }

        let imageDefault = await fetch(defaultGroupPicture);
        let imageBlob = await imageDefault.blob();
        let groupDefault = new File([imageBlob], "default_group.png",
            {
                type: "image/png"
            }
        );

        this.newGroupPhoto = await img_to_base64(groupDefault);

        this.readyToDraw = true;
    },

    watch: {
        selectedOption() {
            // Reset all selected users
            this.usersToAdd = []
            for (let i = 0; i < this.usersGroupFlags.length; i++) {
                this.usersGroupFlags[i] = false;
            }
            this.safeToCreate = false;
        },

        async newGroupName() {
            this.safeToCreate = await this.check_chat_flags();
        },
    }
}
</script>

<template>
<div class="new_chat_dial">
    <div class="new_chat_close_div">
        <div class="new_chat_close_button" role="button" @click="close_dial">
            <img src="../assets/icons/xmark-solid-full.svg" alt="close" class="new_chat_close_img">
        </div>
    </div>
    <h2 class="new_chat_title">Create a new chat</h2>
    <div class="new_chat_selector">
        <button :class="['new_chat_selector_button', { active: selectedOption === 'private' }]"
                @click="selectedOption = 'private'">
            Private Chat
        </button>
        <button :class="['new_chat_selector_button', { active: selectedOption === 'group' }]"
                @click="selectedOption = 'group'">
            Group
        </button>
    </div>
    <div v-if="selectedOption === 'group'" class="new_chat_group_settings">
        <div class="new_chat_group_settings_text">
            <p class="new_chat_text">Group Name</p>
            <input
                v-model="newGroupName"
                type="text"
                placeholder="What's the group name?"
                class="new_chat_input"
            />

            <p class="new_chat_text">Description (optional)</p>
            <textarea
                v-model="newGroupDescription"
                placeholder="What's this group for?"
                class="new_chat_textarea"
                rows="2"
            ></textarea>
        </div>

        <div class="new_chat_group_settings_picture">
            <label class="new_chat_group_picture_label"
                   :style="{ backgroundImage: `url(data:image/jpeg;base64,${newGroupPhoto})` }">
                Click to upload a group picture
                <input class="new_chat_group_picture_button" type="file"
                       accept="image/*" @change="upload_group_picture">
            </label>
        </div>
    </div>

    <div class="new_chat_users_list">
        <div class="extend_full" v-if="selectedOption === 'private' && privateChatUsers.length > 0">
            <UserButton v-for="[index, user] in privateChatUsers.entries()"
                        :key="user.user_id"
                        :user-object="user"
                        :kind-of-chat="selectedOption"
                        :is-selected="usersPrivateFlags[index]"
                        @click="toggle_user_in_conversation(user, index)"></UserButton>
        </div>
        <p v-else-if="selectedOption === 'private' && privateChatUsers.length === 0">Seems like no other users are available to start a private conversation</p>
        <div class="extend_full" v-else-if="selectedOption === 'group' && allUsers.length > 0">
            <UserButton v-for="[index, user] in allUsers.entries()"
                        :key="user.user_id"
                        :user-object="user"
                        :kind-of-chat="selectedOption"
                        :is-selected="usersGroupFlags[index]"
                        @click="toggle_user_in_conversation(user, index)"></UserButton>
        </div>
        <p v-else-if="selectedOption === 'group' && allUsers.length === 0">Seems like no one else uses WASAText yet. Make someone join in order to start a group chat</p>
    </div>

    <div v-if="safeToCreate === true" class="new_chat_button_container">
        <button class="new_chat_confirm_button" @click="create_new_conversation">
            Create the chat
        </button>
    </div>
</div>
</template>

<style scoped>
@import url("../assets/css/new_chat_dial.css");
</style>