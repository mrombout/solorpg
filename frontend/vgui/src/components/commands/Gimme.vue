<template>
  <div class="gimme">
    A shy dog
  </div>
</template>

<script>
export default {
  name: 'Gimme',
  props: {
    text: String
  }
}
</script>

<style scoped>
.gimme {
  border: 1px solid red;
  margin: 2px;
}
</style>
