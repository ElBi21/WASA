<script setup>
import {ref, onMounted} from 'vue';
import TypingBar from "./TypingBar.vue";
import {API_login} from "../services/user-ops";
import ChatTopBar from "./ChatTopBar.vue";

const imageSource = ref('');

onMounted(async () => {
    let base64String = null;
    await API_login("leo").then((result) => {
        base64String = result.userData.profile_pic;
    });

    let format = "jpeg"

    if (!base64String.startsWith('data:image/')) {
        base64String = `data:image/${format};base64,` + base64String;
    }

    console.log(base64String);

    // Set the src attribute of an img element
    imageSource.value = base64String;
})
</script>

<template>
    <div class="chat_window">
        <ChatTopBar></ChatTopBar>
        <img :src="imageSource" alt="Boh" style="display: none;">
        <TypingBar></TypingBar>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_window.css");
</style>
