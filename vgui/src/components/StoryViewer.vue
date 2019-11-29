<template>
  <div class="viewer">
    <component :is="event.type" v-for="event in events" :key="event.type" :text="event.text" />
  </div>
</template>

<script>
import Say from './commands/Say.vue'
import Roll from './commands/Roll.vue'
import Ask from './commands/Ask.vue'
import Gimme from './commands/Gimme.vue'

export default {
  name: 'StoryViewer',
  props: {
    msg: String
  },
  computed: {
    events() {
      return this.$store.state.activeStory.events
    }
  },
  components: {
    Say,
    Roll,
    Ask,
    Gimme
  }
}
</script>

<style scoped>

</style>
