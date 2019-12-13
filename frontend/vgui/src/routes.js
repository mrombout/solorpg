import Vue from 'vue'
import VueRouter from 'vue-router'
import firebase from 'firebase'

import Index from './components/pages/Index.vue'
import Login from './components/pages/Login.vue'
import Editor from './components/pages/Editor.vue'

Vue.use(VueRouter)

const router = new VueRouter({
    mode: 'history',
    routes: [
        {
            path: '/',
            component: Index
        },
        {
            path: '/login',
            component: Login
        },
        {
            path: '/editor',
            component: Editor,
            meta: {
                requiresAuth: true,
            }
        },
        { path: '*', redirect: '/' }
    ]
})

router.beforeEach((to, from, next) => {
    const requiresAuth  = to.matched.some(x => x.meta.requiresAuth)
    const currentUser = firebase.auth().currentUser

    console.log(requiresAuth, currentUser)

    if (requiresAuth && !currentUser) {
        next('/login')
    } else if (requiresAuth && currentUser) {
        next()
    } else {
        next()
    }
})

export default router
