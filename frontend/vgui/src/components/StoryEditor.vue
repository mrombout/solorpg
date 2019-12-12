<template>
  <div class="editor">
    <input type="text" v-model="command" v-on:keyup.enter="onSubmit" />
  </div>
</template>

<script>
export default {
  name: 'StoryEditor',
  data: function() {
    return {
      command: ''
    }
  },
  methods: {
    onSubmit: function() {
      this.$emit('execute-command')
      this.$store.dispatch('addSay', {
        type: 'say',
        text: this.command
      })
      this.command = ''
    }
  }
}
</script>

<style scoped>
.editor textarea {
  width: 800px;
  height: 100px;
}
</style>
