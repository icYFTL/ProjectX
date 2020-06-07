<template>
    <div class="menu-content">
        <ul>
            <li v-for="item in items" :key="item">
                <div class="menu-item">{{item.title}}
                    <hr class="underline">
                </div>
            </li>
        </ul>
        <div class="menu-footer">
            <div class="menu-item">Выход</div>
        </div>

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
                {'title': 'Настройки', 'module': '#'}
            ]
        }
    }
</script>

<style scoped>
    .menu-content {
        background: inherit;
        max-width: 200px;
        height: 93.9vh;
        display: flex;
        flex-flow: column nowrap;

    }
    
    .menu-content ul {
        margin-left: 0;
        padding-left: 0;
        height: 100%;
    }

    .menu-content li {
        margin-top: 10px;
        min-width: 200px;
        min-height: 3em;
        list-style-type: none;
    }

    .menu-footer {
        justify-self: flex-end;
        margin-bottom: 10px;

    }

    .menu-item {
        color: inherit;
        border: none;
        cursor: pointer;
        background: none;
        margin-top: 0.55em;
        -webkit-transition: text-shadow 0.5s ease-in-out;
        -moz-transition: text-shadow 0.5s ease-in-out;
        -o-transition: text-shadow 0.5s ease-in-out;
        transition: text-shadow 0.5s ease-in-out;

    }

    .underline {
        -webkit-transition: width 0.5s ease-in-out;
        -moz-transition: width 0.5s ease-in-out;
        -o-transition: width 0.5s ease-in-out;
        transition: width 0.5s ease-in-out;
        background-color: #979494;
        height: 100%;
        width: 100%;
    }

    .menu-item:hover .underline {
        width: 50%;
    }

    .menu-item:hover {

        text-shadow: 0 0 1px black;
    }


</style>