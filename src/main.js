import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import {BootstrapVue} from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import Vuex from 'vuex'

Vue.use(Vuex)
Vue.use(BootstrapVue)
Vue.config.productionTip = false

const NotFoundComponent = { template: '<div>BAD</div>' }

const router = new VueRouter({
    mode: 'history',
    routes: [
        { path: '*', component: NotFoundComponent }
    ]
});

new Vue({
    render: h => h(App),
    router
}).$mount('#app')
