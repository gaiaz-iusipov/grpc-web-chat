<template>
  <div class="input-group">
    <input
      type="text"
      class="form-control"
      ref="input"
      :maxlength="32"
      v-model.trim="userName"
      v-on:keyup.enter="login"
    >
    <div class="input-group-append">
      <button
        class="btn btn-primary"
        type="button"
        :disabled="isUserNameEmpty"
        @click="login"
      >
        Let's Chat!
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: ''
    }
  },
  mounted() {
    this.focus()
  },
  computed: {
    isUserNameEmpty() {
      return '' === this.userName
    }
  },
  methods: {
    login() {
      this.$emit('submit', this.userName)
    },
    focus() {
      this.$refs.input.focus()
    }
  }
}
</script>
