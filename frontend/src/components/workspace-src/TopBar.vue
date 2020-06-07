<template>
    <div class="top-bar">
        <span>Меню</span>
    </div>
</template>

<script>
    export default {
        name: "TopBar"
    }
</script>

<style scoped>
    .top-bar{
        min-height: 50px;
        background-color: #666666;
        font-size: medium;
        color: white;
        width: 100%;
        overflow: hidden;
    }

    .top-bar span {
        float: left;
        color: #fff;
        text-align: center;
        padding: 14px 16px;
        text-decoration: none;
        font-size: 17px;
        cursor: default;
    }
</style>