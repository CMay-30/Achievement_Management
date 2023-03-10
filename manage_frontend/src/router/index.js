import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/view/Home.vue'
import Login from '@/view/Login.vue'
import Register from '@/view/Register.vue'

Vue.use(Router)

export default new Router({
    routes: [{
            path: '/',
            name: 'Login',
            component: Login
        },
        {
            path: '/Register',
            name: 'Register',
            component: Register
        },
        {
            path: '/Home',
            name: 'Home',
            component: Home
        }
    ]
})