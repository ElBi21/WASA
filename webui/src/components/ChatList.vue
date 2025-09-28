<script>
import SingleChat from "./SingleChat.vue";
import {API_get_conversations} from "../services/user-ops";
import UserPanel from "./UserPanel.vue";

export default {
    props: [
        "userId"
    ],

    data: function() {
        return {
            userChats: [],

            currently_opened_chat: null,
        }
    },

    // Starting point of page
    async mounted() {
        this.userChats = await API_get_conversations("leo");
        console.log(this.userChats)
    },

    methods: {
        async select_chat(event) {
            let chat = event.currentTarget;
            chat.id = "single_chat_selected";

            if (this.currently_opened_chat !== chat && this.currently_opened_chat !== null) {
                this.currently_opened_chat.id = "";
            }

            this.currently_opened_chat = chat;
        }
    },

    components: {
        UserPanel,
        SingleChat
    }
}
</script>

<template>
    <div id="chat_list">
        <div class="chat_list_back">
            <div class="wasatext_bar">
                <h1 class="chat_bar_title">WASAText</h1>
                <div class="home_actions">
                    <button class="home_button">
                        <img class="home_button_icon" id="new_chat_button_icon"
                             src="../assets/icons/comments-solid-full.svg" alt="Settings">
                    </button>
                    <button class="home_button">
                        <img class="home_button_icon" id="settings_button_icon"
                             src="../assets/icons/gear-solid-full.svg" alt="Settings">
                    </button>
                </div>
            </div>
            <div class="chat_list_main">
                <SingleChat v-for="chat in userChats" :chat-id="chat.ID"
                            :chat-name="chat.Name" :last-message-sender="chat.LastSent.sender.display_name"
                            :last-message-body="chat.LastSent.content" :last-message-date="chat.LastSent.timestamp"
                            @click="select_chat">
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