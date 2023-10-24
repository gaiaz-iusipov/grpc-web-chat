<script setup lang="ts">
import { computed, onMounted, ref, defineEmits } from 'vue'

const text = ref<string>('')
const isInputEmpty = computed<boolean>(() => '' === text.value)

const emit = defineEmits(['login'])

function login() {
  emit('login', text.value)
}

const inputRef = ref<HTMLInputElement | null>(null)

onMounted(() => {
  inputRef.value?.focus()
})
</script>

<template>
  <div class="input-group">
    <input
      type="text"
      class="form-control"
      placeholder="Enter your name"
      ref="inputRef"
      :maxlength="32"
      v-model.trim="text"
      v-on:keyup.enter="login"
    >
    <div class="input-group-append">
      <button
        class="btn btn-primary"
        type="button"
        :disabled="isInputEmpty"
        @click="login"
      >
        Let's Chat!
      </button>
    </div>
  </div>
</template>
