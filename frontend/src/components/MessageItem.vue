<script setup lang="ts">
import type { IMessage } from '@/types'

defineProps<{
  message: IMessage
}>()
</script>

<template>
  <div class="media" :class="message.isOur && 'text-end'">
    <div class="media-body text-break">
      <h5 class="mt-0" v-text="message.author.name"/>
      <p class="mb-0" v-text="message.text"/>
    </div>
  </div>
</template>
