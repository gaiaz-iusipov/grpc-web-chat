<template>
  <div class="media" :class="message.isOur && 'text-end'">
    <div class="media-body text-break">
      <h5 class="mt-0" v-text="message.author" />
      <p class="mb-0" v-text="message.text" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    message: {
      type: Object,
      required: true
    }
  }
}
</script>
