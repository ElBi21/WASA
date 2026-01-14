<script>
import {retrieveFromStorage} from "@/services/utils";

export default {
    data: function() {
        return {
            userData: null,
            chatPicture: null,
            chatName: '',
            chatDescription: '',

            privateIndex: 0,
            isThisSelected: false
        }
    },

    async mounted() {
        this.userData = await retrieveFromStorage();
        this.drawButton();
    },

    methods: {
        drawButton() {
            if (this.chatObject.IsPrivate) {
                this.privateIndex = this.chatObject.Users[0].user_id === this.userData.user_id ? 1 : 0;
                this.chatPicture = `data:image/jpeg;base64,` + this.chatObject.Users[this.privateIndex].profile_pic;
                this.chatName = `<b>${this.chatObject.Users[this.privateIndex].display_name}</b> (@${this.chatObject.Users[this.privateIndex].user_id})`;
                this.chatDescription = this.chatObject.Users[this.privateIndex].biography;
            } else {
                this.chatPicture = `data:image/jpeg;base64,` + this.chatObject.Photo;
                this.chatName = `<b>${this.chatObject.Name}</b>`;
                this.chatDescription = this.chatObject.GroupDescription;
            }
        }
    },

    props: [ "chatObject", "isSelected" ],

    watch: {
        chatObject() {
            this.drawButton();
        },

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