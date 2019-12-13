import Vue from 'vue'
import App from './App.vue'

import store from './store'
import router from './routes'

const fb = require('./firebaseConfig.js')

Vue.config.productionTip = false

let app
fb.auth.onAuthStateChanged(() => {
  if (!app) {
    app = new Vue({
      store,
      router,
      render: h => h(App),
    }).$mount('#app')
  }
})
