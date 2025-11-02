<script>
import SingleMessage from "./SingleMessage.vue";
import {API_get_conversation_messages} from "../services/chat-ops";
import {retrieveFromStorage} from "../services/utils";

export default {
    components: { SingleMessage },

    data: function () {
        return {
            userData: null,
            messages: null
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        this.messages = await API_get_conversation_messages(this.chatObj.ID, this.userData.user_id);

        console.log(this.messages);
    },

    props: [ "chatObj" ]
}
</script>

<template>
<div class="chat_messages_container">
    <SingleMessage v-for="message in messages"
       :user-logged="userData.user_id" :message-obj="message"></SingleMessage>
</div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");
</style>