<template>
    <div class="workspace">
        <div class="content">
            <Menu ref="menu" :menu_items_content=menu_items></Menu>
            <ModuleView class="module-view"></ModuleView>
        </div>
    </div>
</template>

<script>
    import Menu from "./workspace-src/menu/Menu";
    import ModuleView from "./workspace-src/ModuleView";

    export default {
        name: "WorkSpace",
        components: {ModuleView, Menu},
        data() {
            return {
                menu_items: [
                    {'title': 'Home', 'module': '#', 'is_active': true},
                    {'title': 'Tasks', 'module': '#', 'is_active': false},
                    {'title': 'Settings', 'module': '#', 'is_active': false}]
            }
        }
    }
</script>

<style lang="scss" scoped>
    $nav-bar-width: 60px;
    $module-view-width: 100vw;
    $module-view-height: 100vh;

    .content {
        display: inline;
    }

    .module-view {
        width: calc(100% - #{$nav-bar-width});
        height: $module-view-height;
        margin-left: $nav-bar-width;
    }

</style>