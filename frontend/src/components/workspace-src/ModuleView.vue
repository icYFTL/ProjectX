<template>
    <div class="module-view">
        <keep-alive>
            <component v-bind:is="module"></component>
        </keep-alive>
    </div>
</template>

<script>
    export default {
        name: "ModuleView",
        props: {
            module: Object
        },
    }
</script>

<style scoped>
    .module-view {
        height: 100%;
        background-color: black;
        align-self: center;
        width: 100%;
    }
</style>