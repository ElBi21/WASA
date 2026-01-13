<script>
import {retrieveFromStorage} from "@/services/utils";
import {API_uncomment_message} from "@/services/comment-ops";

export default {
    data: function () {
        return {
            // loggedUser: null,
            content: null,
            userPFP: null,

            onSelf: false,

            possibleEmojis: {
                'heart suit': '1',
                'thumbs up': '2',
                'face with tears of joy': '3',
                'pouting face': '4',
                'crying face': '5',
                'star': '6'
            }
        }
    },

    emits: [ "removeReaction" ],

    methods: {
        async removeReaction() {
            await API_uncomment_message(this.reactionObj.reaction_id, this.loggedUser.user_id);
            this.$emit("removeReaction");
        }
    },

    async mounted() {
        this.content = this.reactionObj.reaction_content;
        this.userPFP = `data:image/jpeg;base64,` + this.reactionObj.sender.profile_pic;

        this.onSelf = this.loggedUser.user_id === this.messageOwner.user_id;
    },

    name: "SingleReaction",

    props: [ "reactionObj", "messageOwner", "loggedUser" ],
}
</script>

<template>
<div v-if="content !== null" :class="['reaction_button', onSelf ? 'reaction_self' : 'reaction_other' ]"
     role="button" @click="removeReaction">
    <img :src="userPFP" class="reaction_sender_pfp" alt="User that sent the reaction">
    <p class="reaction_content">{{ this.content.FullyQualifiesAs }}</p>
</div>
</template>

<style scoped>
@import url("../assets/css/reactions.css");
</style>