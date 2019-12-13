<template>
  <div class="hello">
    <h1>Editor</h1>
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
import StoryEditor from '../StoryEditor.vue'
import StoryViewer from '../StoryViewer.vue'

export default {
  name: 'Editor',
  props: {
    msg: String
  },
  methods: {
    addRoll() {
      this.$store.dispatch('addRoll')
    },
    addAsk() {

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
