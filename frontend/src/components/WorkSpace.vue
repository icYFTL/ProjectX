<template>
    <div class="workspace">
        <TopBar class="top-bar"></TopBar>
        <Menu class="menu"></Menu>
    </div>
</template>

<script>
    import TopBar from "./workspace-src/TopBar";
    import Menu from "./workspace-src/Menu";
    export default {
        name: "WorkSpace",
        components: {Menu, TopBar}
    }
</script>

<style scoped>
    .workspace {
        display: flex;
        flex-flow: column nowrap;
    }
    .top-bar {
        position: sticky;
        margin-top: -4.5%;
    }
    .menu {
        display: flex;
        flex-flow: row;
        justify-content: flex-start;
    }
</style>