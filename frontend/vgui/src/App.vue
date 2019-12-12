<template>
  <div id="app">
    <StoryViewer />
    <StoryEditor/>
    
    <br />
    <br />
    <br />
    <button @click="addRoll">Roll</button>
    <button @click="addAsk">Ask</button>
  </div>
</template>

<script>
import StoryEditor from './components/StoryEditor.vue'
import StoryViewer from './components/StoryViewer.vue'

export default {
  name: 'app',
  methods: {
    addRoll() {
      this.$store.dispatch('addRoll')
    },
    onExecuteCommand() {
      alert('execute command!')
    }
  },
  components: {
    StoryEditor,
    StoryViewer
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
