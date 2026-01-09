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
            openEditUserDialFlag: false,
            logOutStopTimers: 0,
            chatListCounter: 0,
            refreshUserCounter: 0,
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

        openEditUserDial() {
            this.openEditUserDialFlag = true;
        },

        closeChatDial() {
            this.openNewChatDial = false;
        },

        closeEditUserDial() {
            this.openEditUserDialFlag = false;
        },

        logOutClicked() {
            this.logOutStopTimers = 1;
        },

        refreshChatList() {
            this.chatListCounter += 1;
        },

        async refreshUser() {
            this.refreshUserCounter += 1;
            this.user_data = await retrieveFromStorage();

            console.log(this.user_data);
        }
	},

	async mounted() {
		this.user_data = await retrieveFromStorage();

        console.log(this.user_data);

        this.logOutStopTimers = 0;
	}
}
</script>

<template>
    <Background>
        <div class="main_view">
            <ChatList v-if="this.user_data.user_id !== undefined"
                :refreshCounter="chatListCounter" :refreshUser="refreshUserCounter"
                @chat-selected-emit="handleClickedChatButton"
                @openNewChatDialCL="openChatDial"
                @openEditUserDialCL="openEditUserDial"
                @logOutClicked="logOutClicked"></ChatList>
            <ActivityField :selected-chat-id-prop="chatSelected" :open-chat-dial-external="openNewChatDial"
                :logOutStopTimer="logOutStopTimers" :openEditUserDialExt="openEditUserDialFlag"
                @closeChatDialExternal="closeChatDial"
                @closeEditUserExternal="closeEditUserDial"
                @refreshChatList="refreshChatList"
                @refreshUser="refreshUser"></ActivityField>
        </div>
    </Background>
</template>

<style>
@import url("../assets/css/chat_list.css");
</style>