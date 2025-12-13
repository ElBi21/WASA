<script>
import LoadingSpinner from "../components/LoadingSpinner.vue";
import ChatList from "../components/ChatList.vue";
import Background from "../components/Background.vue";
import ActivityField from "../components/ActivityField.vue";
import {retrieveFromStorage} from "../services/utils";

export default {
    components: {
        ActivityField, Background, ChatList, LoadingSpinner
    },

	data: function() {
		return {
			user_data: {
                user_id: undefined,
                display_name: undefined,
                biography: undefined,
                profile_pic: undefined
            },

            chatSelected: null,
            openNewChatDial: false,
            logOutStopTimers: 0
        }
	},

    // emits: [ "chatSelectedEmit" ],

	methods: {
        handleClickedChatButton(chatID) {
            this.chatSelected = chatID;
        },

        openChatDial() {
            this.openNewChatDial = true;
        },

        closeChatDial() {
            this.openNewChatDial = false;
        },

        logOutClicked() {
            this.logOutStopTimers = 1;
        }
	},

	async mounted() {
		let userData = await retrieveFromStorage();
        Object.assign(this.user_data, userData);

        this.logOutStopTimers = 0;
	}
}
</script>

<template>
    <Background>
        <div class="main_view">
            <ChatList v-if="this.user_data.user_id !== undefined" :user-id='this.user_data.user_id'
                @chat-selected-emit="handleClickedChatButton"
                @openNewChatDialCL="openChatDial"
                @logOutClicked="logOutClicked"></ChatList>
            <ActivityField :selected-chat-id-prop="this.chatSelected" :open-chat-dial-external="openNewChatDial"
                :logOutStopTimer="this.logOutStopTimers" @closeChatDialExternal="closeChatDial"></ActivityField>
        </div>
    </Background>
</template>

<style>
@import url("../assets/css/chat_list.css");
</style>