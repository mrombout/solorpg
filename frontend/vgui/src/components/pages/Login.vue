<template>
  <div class="hello">
    <h1>Login</h1>
    <form>
      <label for="email">Email</label>
      <input type="text" placeholder="you@email.com" id="email" v-model.trim="loginForm.email" />
      <label for="password">Password</label>
      <input type="password" placeholder="********" id="password" v-model.trim="loginForm.password" />
      <button @click="login">Log in</button>
    </form>
  </div>
</template>

<script>
const fb = require('../../firebaseConfig.js')

export default {
  name: 'Login',
  data: function() {
    return {
      loginForm: {
        email: '',
        password: '',
      }
    }
  },
  props: {
    msg: String
  },
  methods: {
    login: function() {
      fb.auth.signInWithEmailAndPassword(this.loginForm.email, this.loginForm.password).then(user => {
        this.$store.commit('setCurrentUser', user)
        this.$store.dispatch('fetchUserProfile')
        this.$router.push('/editor')
      }).catch(err => {
        console.log(err)
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
