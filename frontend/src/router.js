import Router from 'vue-router'
//import store from './store.js'
import Login from './components/Login'
import WorkSpace from "./components/WorkSpace";



let router = new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'app',
            component: WorkSpace,
            meta: {
                requiresAuth: true
            }
        },
        {
            path: '/login',
            name: 'login',
            component: Login
        }
    ]
})

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth))
        //if(store.getters.isLoggedIn)
            next()
        //else
            //next('/login')
    else
        next()

})

export default router