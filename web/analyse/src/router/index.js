import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import NotFound from '../views/NotFound.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        redirect: '/login' // 默认重定向到登录页
    },
    {
        path: '/login',
        name: 'LoginPage',
        component: Login
    },
    {
        path: '/register',
        name: 'RegisterPage',
        component: Register
    },
    {
        path: '/not-found',
        name: 'NotFoundPage',
        component: NotFound
    },
    {
        path: '*', // 匹配所有未定义的路由
        redirect: '/not-found'
    }
]

const router = new VueRouter({
    mode: 'history',
    routes
})

export default router