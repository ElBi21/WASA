<script>
export default {
    data: function() {
        return {
            isThisSelected: false
        }
    },

    props: [ "userObject", "kindOfChat", "isSelected" ],

    watch: {
        kindOfChat() {
            this.isThisSelected = false;
        },

        isSelected() {
            this.isThisSelected = this.isSelected;
        }
    }
}
</script>

<template>
    <div :class="['new_chat_user_button', { active: isThisSelected === true }]" role="button">
        <div>
            <img :src="`data:image/jpeg;base64,` + userObject.profile_pic" alt="User PFP"
            class="new_chat_user_pfp">
        </div>
        <div class="new_chat_user_details">
            <p class="new_chat_user_detail_name text-wrapper"><b>{{ this.userObject.display_name }}</b> (@{{ this.userObject.user_id }})</p>
            <p class="new_chat_user_detail text-wrapper">{{ this.userObject.biography }}</p>
        </div>
    </div>
</template>

<style scoped>
@import url("../assets/css/new_chat_dial.css");
</style>