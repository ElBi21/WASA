<script>
export default {
    props: [ "userLogged", "messageObj" ]
}
</script>

<template>
<div :class="['message_container', { active: messageObj.sender.user_id === userLogged }]">
    <div :class="['message_shape', { active: messageObj.sender.user_id === userLogged }]">
        <p class="message_content">{{ this.messageObj.content }}</p>
    </div>
</div>
</template>

<style scoped>
@import url("../assets/css/chat_messages.css");
</style>