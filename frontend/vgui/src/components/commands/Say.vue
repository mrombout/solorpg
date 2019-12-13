<template>
  <div class="say">
    Say: {{text}}
  </div>
</template>

<script>
export default {
  name: 'Say',
  props: {
    text: String
  },
}
</script>

<style scoped>
.say {
  border: 1px solid green;
  margin: 2px;
}
</style>
