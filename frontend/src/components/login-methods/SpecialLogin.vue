<template>

        <div class="special-login-form">
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Login" v-model="login">
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" placeholder="Password" v-model="password">
                </div>
                <div class="form-group">
                    <label class="form-control">
                        <input type="checkbox" id="save" v-model="save_checked"/>
                        Сохранить пароль?
                    </label>
                </div>
                <button type="submit" class="btn btn-black">Login</button>
        </div>
</template>

<script>
    export default {
        name: "SpecialLogin",
        data() {
            return {
                login: '',
                password: '',
                save_checked: false
            }
        }
    }

</script>

<style scoped>

    .btn-black {
        background-color: #000 !important;
        color: #fff;
        width: 50%;
    }
</style>