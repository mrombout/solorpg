<template>
  <div class="hello">
    <h1>Login</h1>
    <form>
      <transition name="fade">
        <div v-if="performingRequest" class="loading">
          <p>Loading...</p>
        </div>
      </transition>
      <label for="email">Email</label>
      <input type="text" placeholder="you@email.com" id="email" v-model.trim="loginForm.email" />
      <label for="password">Password</label>
      <input type="password" placeholder="********" id="password" v-model.trim="loginForm.password" />
      <button type="button" @click="login">Log in</button>
      <transition name="fade">
        <div v-if="errorMsg != ''" class="error-msg">
          <p>{{ errorMsg }}</p>
        </div>
      </transition>
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
      },
      performingRequest: false,
      errorMSg: '',
    }
  },
  methods: {
    login: function() {
      this.performingRequest = true

      fb.auth.signInWithEmailAndPassword(this.loginForm.email, this.loginForm.password).then(user => {
        this.$store.commit('setCurrentUser', user)
        this.$store.dispatch('fetchUserProfile')
        this.$router.push('/editor')

        this.performingRequest = false
      }).catch(err => {
        console.log(err)
        this.performingRequest = false
        this.errorMsg = err.message
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
