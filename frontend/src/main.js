import Vue from 'vue'
import App from './App.vue'
import {BootstrapVue, IconsPlugin} from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import Vuex from 'vuex'
import Axios from 'axios'
import router from "./router";
import Router from "vue-router";

Vue.use(Vuex)
Vue.use(Axios)
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.config.productionTip = false
Vue.use(Router)

Vue.prototype.$http = Axios;
const token = localStorage.getItem('token')
if (token)
    Vue.prototype.$http.defaults.headers.common['Authorization'] = token



new Vue({
    render: h => h(App),
    router
}).$mount('#app')
