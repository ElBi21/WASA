<script>
import SingleChat from "./SingleChat.vue";
import {API_get_conversations} from "../services/user-ops";
import UserPanel from "./UserPanel.vue";
import {retrieveFromStorage} from "../services/utils";

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
            refresh_timer_interval: 1500
        }
    },

    emits: [ "chatSelectedEmit", "openNewChatDialCL" ],

    methods: {
        async select_chat(chatObj, event) {
            // Set CSS ID
            let chatHTML = event.currentTarget;
            let chatID = chatObj.ID;
            chatHTML.id = "single_chat_selected";

            if (this.currently_opened_chat_html !== chatHTML && this.currently_opened_chat_html !== null) {
                this.currently_opened_chat_html.id = "";
            }

            this.currently_opened_chat_html = chatHTML;
            this.currently_opened_chat_id = chatID;

            // Send value upstream
            this.$emit("chatSelectedEmit", chatObj.ID);
        },

        /*async prepare_chat_for_display(chatID) {

        },*/

        openChatDialEXT() {
            this.$emit("openNewChatDialCL");
        }
    },

    // Starting point of page
    async mounted() {
        await API_get_conversations(this.userId).then((result) => {
            this.userChats = result;
        });

        this.refresh_timer_ID = setInterval(async () => {
            await API_get_conversations(this.userId).then((result) => {
                this.userChats = result;
            });
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
                            :last-message-body="chat.LastSent.content" :last-message-date="chat.LastSent.timestamp"
                            @click="select_chat(chat, $event)">
                </SingleChat>
            </div>
        </div>
        <UserPanel></UserPanel>
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