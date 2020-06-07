<template>
    <div class="workspace">
        <TopBar class="top-bar"></TopBar>
        <div class="content">
        <Menu class="menu-content"></Menu>
        <ModuleView class="module-view">
        </ModuleView>
        </div>
    </div>
</template>

<script>
    import TopBar from "./workspace-src/TopBar";
    import Menu from "./workspace-src/Menu";
    import ModuleView from "./workspace-src/ModuleView";

    export default {
        name: "WorkSpace",
        components: {ModuleView, Menu, TopBar}
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
    .menu-content {
        justify-self: flex-start;
    }

    .content {
        display: flex;
        flex-flow: row nowrap;
    }
    .module-view {
        justify-self: flex-end;
        height: 93.9vh;
        width: 300vh;
    }

</style>