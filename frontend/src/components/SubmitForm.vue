<script setup lang="ts">
import { computed, onMounted, ref, defineEmits } from 'vue'

const text = ref<string>('')
const isTextEmpty = computed<boolean>(() => '' === text.value)

const emit = defineEmits(['submit'])

function submit() {
  if (isTextEmpty.value) {
    return
  }

  emit('submit', text.value)
  text.value = ''
  focus()
}

onMounted(() => {
  focus()
})

const inputRef = ref<HTMLInputElement | null>(null)

function focus() {
  inputRef.value?.focus()
}
</script>

<template>
  <div class="input-group">
    <input
      type="text"
      class="form-control"
      placeholder="Enter your message"
      ref="inputRef"
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
