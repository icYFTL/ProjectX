<template>
    <div class="menu">
        <ul>
            <li v-for="item in items" :key="item">
                <div class="menu-item">{{item.title}}<hr></div>
            </li>
        </ul>

    </div>
</template>

<script>
    export default {
        name: "Menu",
        props: {
            items: Array
        },
        data() {
            this.items = this.items || [
                {'title': 'Задачи', 'module': '#'},
                {'title': 'Настройки', 'module': '#'},
                {'title': 'Выход', 'module': '#'}]
        }
    }
</script>

<style scoped>
    .menu {
        background: inherit;
        max-width: 200px;
        height: 100%;
        display: flex;
        flex-flow: column nowrap;
    }
    .menu ul {
        margin-left: 0;
        padding-left: 0;
        height: 100%;
    }
    .menu li {
        margin-top: 10px;
        min-width: 200px;
        min-height: 3em;
        list-style-type: none;
    }

    .menu-item {
        color: inherit;
        border: none;
        background: none;
        margin-top: 0.55em;
        height: 100%;
        width: 100%;
    }
    .menu-item hr {
        background-color: #979494;

    }
    .menu-item:hover hr {
        -webkit-transition: width 1s ease-in-out;
        -moz-transition: width 1s ease-in-out;
        -o-transition: width 1s ease-in-out;
        transition: width 1s ease-in-out;
        width: 50%;
    }

</style>