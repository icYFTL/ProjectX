<template>
    <div>
        <div class="sidenav">
            <div class="login-main-text">
                <h2>{{app_name}}<br> Login Page</h2>
                <p>{{login_page_description}}</p>
            </div>
        </div>
        <div class="main">
            <div class="btn-group switcher" role="group">
                <button type="button" @click="is_institute_auth = true" class="btn btn-default"
                        v-bind:class="{'btn_switcher-active': is_institute_auth}">Institute
                </button>
                <button type="button" @click="is_institute_auth = false" class="btn btn-default"
                        v-bind:class="{'btn_switcher-active': !is_institute_auth}">Special
                </button>
            </div>
            <InstituteLogin v-if="is_institute_auth" :inst_options=inst_options :group_options=group_options
                            :snp_options=snp_options></InstituteLogin>
            <SpecialLogin v-if="!is_institute_auth"></SpecialLogin>
        </div>
    </div>
</template>

<script>

    import InstituteLogin from "@/components/login-methods/InstituteLogin";
    import SpecialLogin from "@/components/login-methods/SpecialLogin";

    export default {
        name: "Login",
        components: {SpecialLogin, InstituteLogin},
        data() {
            return {
                inst_options: ["None"],
                group_options: ["None"],
                snp_options: ["None"],
                is_institute_auth: false,
                app_name: "Project X",
                login_page_description: "Make world better!"

            }
        },
        methods: {


        }

    }


</script>
<style scoped>
    .switcher {
        margin-top: 40%;
        margin-bottom: 5%;
    }

    @media screen and (max-width: 450px) {
        .switcher {
            margin-top: 10%;
        }

    }

    @media (max-width: 768px) {
        .switcher {
            margin-top: 5%;
        }
    }

    body {
        font-family: "Lato", sans-serif;
    }


    .sidenav {
        height: 100%;
        background-color: #000;
        overflow-x: hidden;
        padding-top: 20px;
    }


    .main {
        display: flex;
        align-items: center;
        flex-direction: column;
        justify-content: center;
        margin-top: -12%;
        padding: 0 10px;
    }

    @media screen and (max-height: 700px) and (max-width: 768px) {
        .sidenav {
            padding-top: -10%;
            margin-top: -10%;
        }

        .main {
            margin-top: 0%;
        }

    }

    .slide-fade-enter-active {
        transition: all .3s ease;
    }

    .slide-fade-leave-active {
        transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
    }

    .btn_switcher-active {
        background-color: #d9fcd9;
        transition: color 3s ease-in-out;
    }


    @media screen and (min-width: 768px) {
        .main {
            margin-left: 40%;
        }

        .sidenav {
            width: 40%;
            position: fixed;
            z-index: 1;
            top: 0;
            left: 0;
        }

    }


    .login-main-text {
        margin-top: 20%;
        padding: 60px;
        color: #fff;
    }

    .login-main-text h2 {
        font-weight: 300;
    }


</style>