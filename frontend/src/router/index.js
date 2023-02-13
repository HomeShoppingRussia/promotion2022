import Vue from 'vue'
import VueRouter from 'vue-router'
import Loader from "vue-ui-preloader";

Vue.use(VueRouter)
Vue.use(Loader);

const routes = [
    {
        path: '/public/phone.png',
        component: Image,
        props: true,
    },
    {
        path: '/',
        name: 'home',
        component: Vue
    },
];

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router