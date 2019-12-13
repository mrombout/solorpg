<template>
  <div id="app">
    <TitleBar></TitleBar>
    <router-view></router-view>
  </div>
</template>

<script>
import TitleBar from './components/TitleBar'

export default {
  name: 'app',
  components: {
    TitleBar
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
