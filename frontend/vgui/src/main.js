import Vue from 'vue'
import App from './App.vue'
import store from './store'
import VueRouter from 'vue-router'
const fb = require('./firebaseConfig.js')

Vue.config.productionTip = false

import Index from './components/pages/Index.vue'
import Login from './components/pages/Login.vue'
import Editor from './components/pages/Editor.vue'

const routes = [
  { path: '/', component: Index },
  { path: '/login', component: Login },
  { path: '/editor', component: Editor }
]

const router = new VueRouter({
  routes: routes
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

Vue.use(VueRouter)

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
