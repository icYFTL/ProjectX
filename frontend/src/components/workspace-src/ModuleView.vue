<template>
    <div class="module-view" @mouseover="secret" @mouseleave="recover_secret">
        <h1 style="color: white;"  >{{secret_text}}</h1>
    </div>
</template>

<script>
    export default {
        name: "ModuleView",
        methods: {
            secret: function () {

                this.secret_text = "      well  rag  s   ";
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