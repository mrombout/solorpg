<template>
  <div class="roll">
    {{result.Result}} = <span class="dice">⚀ ⚁ ⚂ ⚃ ⚄ ⚅</span>
  </div>
</template>

<script>
export default {
  name: 'Roll',
  props: {
    result: Object,
  }
}
</script>

<style scoped>
.roll {
  font-size: 21px;
}

.roll .dice {
  font-size: 32px;
}
</style>
