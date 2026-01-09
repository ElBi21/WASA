<script>
import {img_to_base64, retrieveFromStorage} from "@/services/utils";
import {API_set_group_photo, API_set_group_description, API_set_group_name} from "@/services/chat-ops";

export default {
    data: function() {
        return {
            userID: null,

            groupName: "",
            groupDescription: "",
            groupPhoto: "",
        }
    },

    emits: [ "closeEditGroupDial" ],

    methods: {
        async upload_group_picture(event) {
            let picture = event.target.files[0];
            this.groupPhoto = await img_to_base64(picture);
        },

        async update_chat() {
            await API_set_group_name(this.chatObj.ID, this.groupName, this.userID);
            await API_set_group_description(this.chatObj.ID, this.groupDescription, this.userID);
            await API_set_group_photo(this.chatObj.ID, this.groupPhoto, this.userID);

            this.$emit("closeEditGroupDial");
        },

        async close_dial() {
            this.$emit("closeEditGroupDial");
        }
    },

    async mounted() {
        console.log(this.chatObj);

        let userData = await retrieveFromStorage();
        this.userID = userData.user_id;

        this.groupName = this.chatObj.Name;
        this.groupDescription = this.chatObj.GroupDescription;
        this.groupPhoto = this.chatObj.Photo;
    },

    name: "EditGroupDial",

    props: [ "chatObj" ],
}
</script>

<template>
    <div class="edit_group_dial">
        <div class="edit_group_close_div">
            <div class="edit_group_close_button" role="button" @click="close_dial">
                <img src="../assets/icons/xmark-solid-full.svg" alt="close" class="edit_group_close_img">
            </div>
        </div>
        <h2 class="edit_group_title">Edit the group settings</h2>
        <div class="edit_group_settings">
            <div class="edit_group_settings_text">
                <p class="edit_group_text">Group Name</p>
                <input
                    v-model="groupName"
                    type="text"
                    placeholder="What's the group name?"
                    class="edit_group_input"
                />

                <p class="edit_group_text">Description (optional)</p>
                <textarea
                    v-model="groupDescription"
                    placeholder="What's this group for?"
                    class="edit_group_textarea"
                    rows="2"
                ></textarea>
            </div>

            <div class="edit_group_settings_picture">
                <label class="edit_group_picture_label"
                       :style="{ backgroundImage: `url(data:image/jpeg;base64,${groupPhoto})` }">
                    Click to upload a group picture
                    <input class="edit_group_picture_button" type="file"
                           accept="image/*" @change="upload_group_picture">
                </label>
            </div>
        </div>

        <div v-if="groupName.length >= 1" class="edit_group_button_container">
            <button class="edit_group_confirm_button" @click="update_chat">
                Update the group
            </button>
        </div>
    </div>
</template>

<style scoped>
@import url("../assets/css/edit_dials.css");
</style>