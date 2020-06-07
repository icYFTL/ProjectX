<template>
    <div class="col-md-6 col-sm-12">
        <div class="special-login-form">
            <form @submit.prevent="auth">
                <div class="form-group">
                    <input required type="text" class="form-control" placeholder="Login" v-model="login">
                </div>
                <div class="form-group">
                    <input required type="password" class="form-control" placeholder="Password" v-model="password">
                </div>
                <div class="form-group">
                    <label class="form-control">
                        <input type="checkbox" id="save" v-model="save_checked"/>
                        Сохранить пароль?
                    </label>
                </div>
                <button type="submit" v-on:click="auth" class="btn btn-black">Login</button>
            </form>
        </div>
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
        },
        methods: {
            auth: function () {
                this.$router.push('/');
                // this.$store.dispatch('login', {
                //     'login': this.login,
                //     'password': this.password
                // })
                //     .then(() => this.$router.push('/'))
                //     .catch(err => console.log(err))
            }
        }
    }

</script>

<style scoped>
    .btn-black {
        background-color: #000 !important;
        color: #fff !important;
        width: 50%;
    }

</style>