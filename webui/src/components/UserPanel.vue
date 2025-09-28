<script setup>
import {ref, onMounted} from 'vue';
import {API_login} from "../services/user-ops";

const imageSource = ref('');
const userData = ref('');

onMounted(async () => {
    let base64String = null;
    await API_login("leo").then((result) => {
        base64String = result.userData.profile_pic;
        userData.value = result.userData;
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
    <div class="user_panel">
        <img :src="imageSource" alt="User PFP" class="user_pfp">
        <div class="user_data">
            <p class="user_info"><b>{{ userData.display_name }}</b> &nbsp(@{{ userData.user_id }})</p>
            <p class="user_bio">{{ userData.biography }}</p>
        </div>
        <button class="home_button" id="logout_button">
            <img class="home_button_icon" id="logout_button_icon"
                 src="../assets/icons/door-open-solid-full.svg" alt="Settings">
        </button>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_list.css");
</style>