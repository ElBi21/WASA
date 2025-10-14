<script setup>
import {ref, onMounted} from 'vue';

const imageSource = ref('');
const userData = ref('');

onMounted(async () => {
    let user_data = {};
    let userDataFromLogin = JSON.parse(localStorage.getItem("user"));
    Object.assign(user_data, userDataFromLogin);

    userData.value = user_data;

    let format = "jpeg";

    if (!user_data.profile_pic.startsWith('data:image/')) {
        user_data.profile_pic = `data:image/${format};base64,` + user_data.profile_pic;
    }

    // Set the src attribute of an img element
    imageSource.value = user_data.profile_pic;
})
</script>

<script>
export default {
    methods: {
        async goToLogin() {
            this.$router.push({ path: "/", replace: true });
        }
    }
}
</script>

<template>
    <div class="user_panel">
        <img :src="imageSource" alt="User PFP" class="user_pfp">
        <div class="user_data">
            <p class="user_info"><b>{{ userData.display_name }}</b>&nbsp(@{{ userData.user_id }})</p>
            <p class="user_bio">{{ userData.biography }}</p>
        </div>
        <button class="home_button" id="logout_button" @click="goToLogin">
            <img class="home_button_icon" id="logout_button_icon"
                 src="../assets/icons/door-open-solid-full.svg" alt="Settings">
        </button>
    </div>
</template>

<style scoped>
@import url("../assets/css/chat_list.css");
</style>