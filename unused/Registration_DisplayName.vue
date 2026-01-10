<script>
import text_json from "../webui/src/assets/texts/login.json";

export default {
    data: function () {
        return {
            username: "",
            newDisplayName: "",

            userData: {},
            textJson: text_json,

            current_register_step: 0
        }
    },

    methods: {

    }
}
</script>

<template>
    <h1 id="register_title">Welcome to WASAText</h1>
    <p id="register_description">Hi&nbsp;<b>@{{ user_data.user_id }}</b>, do you want to set a new display name?</p>
    <div id="new_display_form">
        <div class="form_zone">
            <input v-model="newDisplayName" class="login_register_input" placeholder="Choose a display name" v-on:keyup.enter="set_new_display_name">
            <button id="login_button" @click="set_new_display_name">
                <img id="login_arrow" src="../webui/src/assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
            </button>
        </div>
        <button class="skip_step" @click="skip_register_step">I don't want to set a new display name</button>
    </div>
    <div id="progress_bars">
        <div class="progress" id="register_dispname"></div>
        <div class="progress" id="register_bio"></div>
        <div class="progress" id="register_photo"></div>
    </div>
</template>

<style scoped>

</style>