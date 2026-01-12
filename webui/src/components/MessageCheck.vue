<script>
export default {
    name: "MessageCheck",

    data: function() {
        return {
            allReceived: false,
            allSeen: false
        }
    },

    methods: {
        countSeen() {
            this.allSeen = this.seenList.length === this.userList.length;
        }
    },

    async mounted() {
        this.countSeen();
    },

    props: [ "recvList", "seenList", "userList" ],

    watch: {
        seenList() {
            this.countSeen();
        }
    }
}
</script>

<template>
    <div class="message_status">
        <img v-if="allSeen === false" class="message_status_icon" src="../assets/icons/single-check-full.svg" alt="All users received this message">
        <img v-else class="message_status_icon" src="../assets/icons/double-check-full.svg" alt="All users have seen this message">
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");
</style>