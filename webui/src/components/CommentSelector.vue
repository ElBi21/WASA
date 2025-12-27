<script>
import {API_comment_message, API_uncomment_message} from "@/services/comment-ops";

export default {
    data: function () {
        return {

        }
    },

    emits: [ 'closeSelector' ],

    methods: {
        async comment(reaction) {
            // Not really necessary, API_comment_message can remove any present reaction, but it doesn't refresh!
            if (this.hasUserReacted !== undefined && this.hasUserReacted.flag) {
                await API_uncomment_message(this.hasUserReacted.reaction_id, this.commentingUserID);
            }

            await API_comment_message(this.messageID, this.commentingUserID, reaction);
            this.$emit("closeSelector");
        }
    },

    name: "CommentSelector",

    props: [ 'showSelector', 'messageID', 'commentingUserID', 'hasUserReacted' ]
}
</script>

<template>
<div v-if="showSelector" class="react_dial">
    <div class="react_emoji" role="button" @click="comment('♥️')">♥️</div>
    <div class="react_emoji" role="button" @click="comment('👍')">👍</div>
    <div class="react_emoji" role="button" @click="comment('😂')">😂</div>
    <div class="react_emoji" role="button" @click="comment('😡')">😡</div>
    <div class="react_emoji" role="button" @click="comment('😢')">😢</div>
    <div class="react_emoji" role="button" @click="comment('⭐')">⭐</div>
</div>
</template>

<style scoped>
.react_dial {
    width: fit-content;
    height: 40px;

    padding: 0 3px 0 3px;
    margin: 4px 0 0 0;

    background-color: white;
    border-radius: 20px;
    box-shadow: 0 0 30px rgba(0, 0, 0, 0.5);

    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;

    z-index: 10;
}

.react_emoji {
    width: 32px;
    height: 32px;

    margin: 4px 1px 4px 1px;

    border-radius: 16px;
    background-color: #eaeaed;

    transition: background-color 0.2s ease-in-out;

    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;

    font-size: 16px;
}

.react_emoji:hover {
    background-color: #d9d9de;
    transition: background-color 0.2s ease-in-out;
}
</style>