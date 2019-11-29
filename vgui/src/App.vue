<template>
  <div id="app">
    <StoryViewer />
    <StoryEditor/>
    
    <br />
    <br />
    <br />
    Clicked: {{ $store.state.count }} times, count is {{ evenOrOdd }}
    <button @click="increment">+</button>
    <button @click="decrement">-</button>
    <button @click="incrementIfOdd">Increment if odd</button>
    <button @click="incrementAsync">Increment async</button>
  </div>
</template>

<script>
import StoryEditor from './components/StoryEditor.vue'
import StoryViewer from './components/StoryViewer.vue'
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'app',
  computed: mapGetters([
    'evenOrOdd'
  ]),
  methods: {
    ...mapActions([
      'increment',
      'decrement',
      'incrementIfOdd',
      'incrementAsync'
    ]),
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
