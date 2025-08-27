<script>
export default {
    data: function () {
        return {
            username: "",
            user_data: {}
        }
    },
    methods: {
        async login() {
            try {
                let response = await this.$axios.post("/session", {
                    headers: {
                        "Content-Type": "application/json",
                    },
                    name: "leo"
                });

                this.user_data = response.data;

                console.log(this.user_data);
            } catch (error) {
                console.log(error);
            }
        }
    }
}
</script>

<template>
    <div id="login_bg_1">
        <div id="login_bg_2">
            <div id="login_body">
                <div id="login_main_form">
                    <h1 id="login_title">WASAText</h1>
                    <div id="login_form">
                        <input id="login_input" placeholder="Insert your username">
                        <button id="login_button" @click="login">
                            <img id="login_arrow" src="../assets/icons/arrow-right-solid-full.svg" alt="Login arrow">
                        </button>
                    </div>
                    <p>{{ user_data.Biography }}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<style>

</style>