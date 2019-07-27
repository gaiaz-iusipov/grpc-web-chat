<template>
  <div class="input-group">
    <input
      type="text"
      class="form-control"
      placeholder="Enter your message"
      ref="input"
      v-model.trim="text"
      v-on:keyup.enter="submit"
    >
    <div class="input-group-append">
      <button
        class="btn btn-success"
        type="button"
        :disabled="isTextEmpty"
        @click="submit"
      >
        Submit
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      text: ''
    }
  },
  mounted() {
    this.focus()
  },
  computed: {
    isTextEmpty() {
      return '' === this.text
    }
  },
  methods: {
    submit() {
      if (0 === this.text.length) {
        return
      }

      this.$emit('submit', this.text)
      this.text = ''
      this.focus()
    },
    focus() {
      this.$refs.input.focus()
    }
  }
}
</script>
