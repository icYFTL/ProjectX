<template>
    <div class="module-view" >
        <h1 style="color: white; cursor: default;" @mouseover="secret" @mouseleave="recover_secret" >{{secret_text}}</h1>
    </div>
</template>

<script>
    export default {
        name: "ModuleView",
        methods: {
            secret: function () {
                this.secret_text = "We don't always choose the right way";
            },
            recover_secret: function () {
                this.secret_text = "Here dwell dragons...";
            }
        },
        data: function () {
            return {
                secret_text: "Here dwell dragons..."
            }
        }
    }
</script>

<style scoped>
    .module-view {
        height: 100%;
        background-color: black;
        align-self: center;
        width: 100%;
    }

    .module-view h1 {

    }
</style>