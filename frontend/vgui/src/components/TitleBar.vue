<template>
  <div class="titlebar">
    Titlebar
  </div>
</template>

<script>
export default {
  name: 'TitleBar',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.titlebar {
  background-color: #330033;
}
</style>
