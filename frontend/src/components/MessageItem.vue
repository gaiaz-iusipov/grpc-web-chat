<script setup lang="ts">
import { defineProps } from 'vue'
import type { IMessage } from '@/types'

defineProps<{
  message: IMessage,
  showAuthor: boolean,
}>()
</script>

<template>
  <div class="media" :class="message.isOur && 'text-end'">
    <div class="media-body text-break">
      <h5
        v-if="showAuthor"
        v-text="message.author.name"
      />
      <p class="mb-0" v-text="message.text"/>
    </div>
  </div>
</template>
