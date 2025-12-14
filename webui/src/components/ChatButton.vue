<script>
export default {
    data: function() {
        return {
            chatPicture: null,
            chatName: '',
            chatDescription: '',

            isThisSelected: false
        }
    },

    mounted() {
        if (this.chatObject.IsPrivate) {
            this.chatPicture = `data:image/jpeg;base64,` + this.chatObject.Users[0].profile_pic;
            this.chatName = `<b>${this.chatObject.Users[0].display_name}</b> (@${this.chatObject.Users[0].user_id})`;
            this.chatDescription = this.chatObject.Users[0].biography;
        } else {
            this.chatPicture = `data:image/jpeg;base64,` + this.chatObject.Photo;
            this.chatName = `<b>${this.chatObject.Name}</b>`;
            this.chatDescription = this.chatObject.GroupDescription;
        }
    },

    props: [ "chatObject", "isSelected" ],

    watch: {
        isSelected() {
            this.isThisSelected = this.isSelected;
        }
    }
}
</script>

<template>
    <div v-if="chatObject !== undefined" :class="['forward_msg_user_button', { active: isThisSelected === true }]" role="button">
        <div>
            <img :src="this.chatPicture" alt="Chat picture"
            class="forward_msg_user_pfp">
        </div>
        <div class="forward_msg_user_details">
            <p class="forward_msg_user_detail_name text-wrapper" v-html="chatName"></p>
            <p class="forward_msg_user_detail text-wrapper">{{ this.chatDescription }}</p>
        </div>
    </div>
</template>

<style scoped>
@import url("../assets/css/forward_msg_dial.css");
</style>