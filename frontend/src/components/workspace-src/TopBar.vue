<template>
    <div class="top-bar">
        <div v-on:click="toggleMenuRef1.toggle_hide">
            <b-icon icon="chevron-double-left"></b-icon>
            <span>Меню</span>
        </div>
    </div>
</template>

<script>
    export default {
        name: "TopBar",
        props: {
          toggleMenuRef: Object
        },
        data() {
            return {
                toggleMenuRef1: this.toggleMenuRef1 || Object
            }
        }
    }
</script>

<style scoped>
    .top-bar {
        min-height: 50px;
        background-color: #666666;
        font-size: medium;
        color: white;
        width: 100%;
        overflow: hidden;
    }

    .top-bar div {
        float: left;
        color: #fff;
        text-align: center;
        padding: 14px 16px;
        text-decoration: none;
        font-size: 17px;
        cursor: pointer;
    }
</style>