<template>
  <div id="app">
    <TitleBar>
      VGUI <a v-if="currentUser" @click="logout" href="#">| Logout</a>
    </TitleBar>
    <router-view></router-view>
  </div>
</template>

<script>
import TitleBar from './components/TitleBar'
import { mapState } from 'vuex'
const fb = require('./firebaseConfig.js')

export default {
  name: 'app',
  components: {
    TitleBar
  },
  computed: {
    ...mapState(['currentUser'])
  },
  methods: {
    logout() {
      fb.auth.signOut().then(() => {
          this.$store.dispatch('clearData')
          this.$router.push('/login')
      }).catch(err => {
        console.log(err)
      })
    }
  }
}
</script>

<style>
body {
  padding: 0;
  margin: 0;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding: 0;
  margin: 0;
}
</style>
