<template>
  <div class="viewer">
    <Say text="Hiya" />
    <component :is="event.type" v-for="event in events" :key="event.id + '_' + event.type" :text="event.text" :result="event" />
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
