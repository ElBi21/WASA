<script>
import MessageCheck from "@/components/MessageCheck.vue";
import {API_delete_message, API_retrieve_message} from "@/services/message-ops";
import CommentSelector from "@/components/CommentSelector.vue";
import SingleReaction from "@/components/SingleReaction.vue";

export default {
    components: {SingleReaction, CommentSelector, MessageCheck},

    data: function() {
        return {
            messageFormattedDate: null,
            senderPFP: null,

            openCommentSelector: false,

            repliedMessage: null,
            replyToText: '',

            messageReactions: null,

            hasUserReacted: {
                flag: false,
                reaction_id: null
            }
        }
    },

    emits: [ "refreshChat", "openForwardDial", "startReplyToMessage" ],

    methods: {
        async deleteMessage() {
            await API_delete_message(this.messageObj.message_id, this.userLogged.user_id);
            this.$emit("refreshChat");
        },

        openForwardDial() {
            this.$emit("openForwardDial", this.messageObj);
        },

        startReply() {
            this.$emit("startReplyToMessage", this.messageObj);
        },

        async closeCommentSelector() {
            this.openCommentSelector = false;
            this.messageReactions = null;
            this.$emit("refreshChat");
        },
    },

    async mounted() {
        let messageTimestamp = new Date(Date.parse(this.messageObj.timestamp));
        this.messageFormattedDate = `${messageTimestamp.getHours().toString().padStart(2, '0')}:${messageTimestamp.getMinutes().toString().padStart(2, '0')}`;

        this.senderPFP = `data:image/jpeg;base64,` + this.messageObj.sender.profile_pic;

        if (this.messageObj.replying !== 0) {
            this.repliedMessage = await API_retrieve_message(this.messageObj.replying, this.userLogged.user_id);
        }

        this.messageReactions = this.reactionsObj;

        if (this.reactionsObj !== null) {
            let userReaction = this.reactionsObj.findIndex(reaction => reaction.sender.user_id === this.userLogged.user_id);
            this.hasUserReacted.flag = userReaction !== -1;

            if (this.hasUserReacted.flag) {
                this.hasUserReacted.reaction_id = this.reactionsObj[userReaction].reaction_id;
            }
        }

        // console.log(this.messageObj);
    },

    props: [ "userLogged", "messageObj", "isChatPrivate", "chatUsers", "refreshEnforcer", "reactionsObj" ],

    watch: {
        reactionsObj() {
            this.messageReactions = this.reactionsObj;
        }
    }
}
</script>

<template>
<div :class="['message_container', { active: messageObj.sender.user_id === userLogged.user_id }]">
    <div class="message_side_container message_side_logged_user" v-if="messageObj.sender.user_id === userLogged.user_id">
        <div class="message_side_button msg_btn_top" role="button" @click="deleteMessage" v-if="!this.messageObj.deleted">
            <img src="../assets/icons/trash-can-solid-full.svg" class="msg_side_icon" alt="Edit the message">
        </div>
        <div class="message_side_button" role="button" v-if="!this.messageObj.deleted" @click="openForwardDial">
            <img src="../assets/icons/arrows-turn-right-solid-full.svg" class="msg_side_icon" alt="Forward the message">
        </div>
        <div :class="['message_side_button', this.messageObj.deleted ? 'msg_btn_round' : 'msg_btn_bottom']"
             role="button" @click="openCommentSelector = !openCommentSelector">
            <img src="../assets/icons/heart-circle-plus-solid-full.svg" class="msg_side_icon" alt="React to the message">
        </div>
        <CommentSelector :showSelector="openCommentSelector" @closeSelector="closeCommentSelector"
                         :messageID="messageObj.message_id" :commentingUserID="userLogged.user_id" :hasUserReacted="hasUserReacted"
                         :position="messageObj.sender.user_id !== userLogged.user_id ? 'right' : 'left'"></CommentSelector>
    </div>
    <div :class="['message_shape', { active: messageObj.sender.user_id === userLogged.user_id }]">
        <div class="message_sender_info" v-if="!isChatPrivate">
            <img class="message_sender_photo" v-if="messageObj.sender.user_id !== userLogged.user_id"
                 :src="this.senderPFP" alt="User PFP">
            <p class="message_sender_display_name" v-if="messageObj.sender.user_id !== userLogged.user_id">{{ this.messageObj.sender.display_name }}</p>
        </div>

        <div class="message_picture" v-if="messageObj.photo !== ''">
            <img class="message_image" :src="`data:image/jpeg;base64,` + messageObj.photo" alt="The message's photo">
        </div>

        <div class="reply_preview_container" v-if="repliedMessage !== null">
            <img src="../assets/icons/comment-dots-regular-full.svg" class="reply_icon" alt="Original message">
            <p class="reply_preview_content">
                {{ repliedMessage.deleted ? 'This message has been deleted' : repliedMessage.content }}</p>
        </div>

        <p v-if="!this.messageObj.deleted" class="message_content">{{ this.messageObj.content }}</p>
        <div v-else class="message_deleted">
            <img src="../assets/icons/ban-solid-full.svg" class="message_deleted_icon">
            <p class="message_deleted_text">This message has been deleted</p>
        </div>

        <div v-if="messageReactions !== null && this.userLogged !== null" class="reactions_container">
            <SingleReaction v-for="reaction of messageReactions" :key="reaction.reaction_id"
                            @removeReaction="this.$emit('refreshChat')" :loggedUser="this.userLogged"
                            :reactionObj="reaction" :messageOwner="messageObj.sender"></SingleReaction>
        </div>

        <div class="message_metadata">
            <p class="message_time">{{ this.messageFormattedDate }}</p>
            <MessageCheck v-if="messageObj.sender.user_id === userLogged.user_id" :recvList="messageObj.received"
                          :seenList="messageObj.seen" :userList="chatUsers"></MessageCheck>
            <img v-if="messageObj.forwarded" src="../assets/icons/arrows-turn-right-solid-full.svg" class="forwarded_icon" alt="Forwarded message">
        </div>
    </div>
    <div class="message_side_container" v-if="messageObj.sender.user_id !== userLogged.user_id">
        <div class="message_side_button msg_btn_top" role="button" v-if="!this.messageObj.deleted" @click="openForwardDial">
            <img src="../assets/icons/arrows-turn-right-solid-full.svg" class="msg_side_icon" alt="Forward the message">
        </div>
        <div class="message_side_button" role="button" v-if="!this.messageObj.deleted" @click="startReply">
            <img src="../assets/icons/comment-dots-regular-full.svg" class="msg_side_icon" alt="Reply to the message">
        </div>
        <div :class="['message_side_button', this.messageObj.deleted ? 'msg_btn_round' : 'msg_btn_bottom']"
             role="button" @click="openCommentSelector = !openCommentSelector">
            <img src="../assets/icons/heart-circle-plus-solid-full.svg" class="msg_side_icon" alt="React to the message">
        </div>
        <CommentSelector :showSelector="openCommentSelector" @closeSelector="closeCommentSelector"
            :messageID="messageObj.message_id" :commentingUserID="userLogged.user_id" :hasUserReacted="hasUserReacted"
            :position="messageObj.sender.user_id !== userLogged.user_id ? 'right' : 'left'"></CommentSelector>
    </div>

</div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");
</style>