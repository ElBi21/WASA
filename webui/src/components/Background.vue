<script>
export default {
    props: [
        "backgroundStyle"
    ],
}
</script>

<template>
    <div class="bg_general" id="bg_1">
        <div class="bg_general" id="bg_2">
            <div class="bg_general" id="bg_3">
                <slot></slot>
            </div>
        </div>
    </div>
</template>

<style>
@import url("../assets/css/background.css");
</style>