<script>
import LoadingSpinner from "../components/LoadingSpinner.vue";
import ChatList from "../components/ChatList.vue";
import Background from "../components/Background.vue";
import ActivityField from "../components/ActivityField.vue";

export default {
    components: {ActivityField, Background, ChatList, LoadingSpinner},
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		// this.refresh()
	}
}
</script>

<template>
    <Background>
        <div class="main_view">
            <ChatList :user-id="'Leo'"></ChatList>
            <ActivityField></ActivityField>
        </div>
    </Background>
</template>

<style>
@import url("../assets/css/chat_list.css");
</style>