<script>
import SingleChat from "./SingleChat.vue";
import UserPanel from "./UserPanel.vue";
import {API_get_conversations} from "../services/user-ops";

export default {
    async beforeDestroy() {
        if (this.refresh_timer_ID) {
            clearInterval(this.refresh_timer_ID);
            this.refresh_timer_ID = null;
        }
    },

    components: {SingleChat, UserPanel},

    data: function() {
        return {
            userChats: [],

            currently_opened_chat_id: null,
            currently_opened_chat_html: null,

            refresh_timer_ID: null,
            refresh_timer_interval: 1500,

            stopReloading: true,
        }
    },

    emits: [ "chatSelectedEmit", "openNewChatDialCL", "logOutClicked" ],

    methods: {
        async select_chat(chatObj) {
            this.currently_opened_chat_id = chatObj.ID;

            // Send value upstream
            this.$emit("chatSelectedEmit", chatObj.ID);
        },

        openChatDialEXT() {
            this.$emit("openNewChatDialCL");
        },

        logOutHandler() {
            this.$emit('logOutClicked', 1);
            this.stopReloading = 1;
        }
    },

    // Starting point of page
    async mounted() {
        this.stopReloading = false;
        await API_get_conversations(this.userId).then((result) => {
            this.userChats = result;
        });

        this.refresh_timer_ID = setInterval(async () => {
            if (this.stopReloading) {
                this.stopReloading = !this.stopReloading;
                clearInterval(this.refresh_timer_ID);
                this.refresh_timer_interval = null;
            } else {
                await API_get_conversations(this.userId).then((result) => {
                    this.userChats = result;
                });
            }
        }, this.refresh_timer_interval);
    },

    props: [
        "userId"
    ],
}
</script>

<template>
    <div id="chat_list">
        <div class="chat_list_back">
            <div class="wasatext_bar">
                <h1 class="chat_bar_title">WASAText</h1>
                <div class="home_actions">
                    <button class="home_button" role="button" @click="openChatDialEXT">
                        <img class="home_button_icon" id="new_chat_button_icon"
                             src="../assets/icons/comments-solid-full.svg" alt="Settings">
                    </button>
                    <button class="home_button" role="button">
                        <img class="home_button_icon" id="settings_button_icon"
                             src="../assets/icons/gear-solid-full.svg" alt="Settings">
                    </button>
                </div>
            </div>
            <div class="chat_list_main">
                <SingleChat v-for="chat in userChats" :chat-id="chat.ID"
                            :last-message-sender="chat.LastSent.sender.display_name"
                            :last-message-body="chat.LastSent.deleted ? 'Deleted message' : chat.LastSent.content"
                            :last-message-date="chat.LastSent.timestamp"
                            :isChatSelected="chat.ID === currently_opened_chat_id"
                            @click="select_chat(chat)">
                </SingleChat>
            </div>
        </div>
        <UserPanel @logOutClicked="logOutHandler"></UserPanel>
    </div>
</template>

<style scoped>
#chat_list {
    width: 27%;

    display: flex;
    flex-direction: column;
    justify-content: start;
    align-content: center;
}
</style>