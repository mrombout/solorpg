import Vue from 'vue'
import VueRouter from 'vue-router'

import Index from './components/pages/Index.vue'
import Login from './components/pages/Login.vue'
import Editor from './components/pages/Editor.vue'

const fb = require('./firebaseConfig.js')

Vue.use(VueRouter)

const router = new VueRouter({
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
        }
    ]
})

router.beforeEach((to, from, next) => {
    if (to.path == '/login') {
        next()
        return
    }

    if (fb.auth.currentUser == null) {
        next('/login')
        return
    }

    next()
})

export default router
