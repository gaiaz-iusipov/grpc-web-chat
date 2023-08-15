<script setup lang="ts">
import { computed, nextTick, onMounted, ref } from 'vue'
import { v4 as uuidv4 } from 'uuid'
import type { IMessage, IUser } from '@/types'
import type { Message } from '@/chat/v1/chat'
import Client from '@/client'
import SubmitForm from '@/components/SubmitForm.vue'
import LoginForm from '@/components/LoginForm.vue'
import MessageItem from '@/components/MessageItem.vue'

const client = new Client(import.meta.env.VITE_API_URL)

const user = ref<IUser>({
  id: uuidv4(),
  name: ''
})

const isLoggedIn = computed<boolean>(() => '' !== user.value.name)

const messages = ref<IMessage[]>([])

onMounted(() => {
  client.subscribe(user.value.id, async (message: Message) => {
    if (!message.client) {
      return
    }

    const user: IUser = {
      id: message.client.uuid,
      name: message.client.name
    }
    await addMessage(user, message.text, false)
  })
})

async function onSubmit(text: string) {
  await client.sendMessage(user.value, text)
  await addMessage(user.value, text, true)
}

function onLogin(text: string) {
  user.value.name = text
}

const messagesBottomRef = ref<Element | null>(null)

async function addMessage(user: IUser, text: string, isOur: boolean) {
  messages.value.push({
    author: user,
    text: text,
    isOur: isOur,
    createdAt: new Date()
  })

  await nextTick()
  // Code that will run only after the entire view has been re-rendered
  messagesBottomRef.value?.scrollIntoView({ behavior: 'smooth' })
}
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col col-md-6 offset-md-3">
        <div class="card vh-100">
          <div class="card-header">
            <img alt="gRPC" src="@/assets/grpc.svg"> Web Chat
          </div>
          <div class="card-body overflow-y-scroll">
            <MessageItem
              v-for="(message, index) in messages"
              :key="index"
              :message="message"
              :show-author="!index || messages[index-1].author.id !== message.author.id"
            />
            <p class="card-text" v-if="!messages.length">No Messages yet</p>
            <div ref="messagesBottomRef"></div>
          </div>
          <div class="card-footer">
            <SubmitForm @submit="onSubmit" v-if="isLoggedIn"/>
            <LoginForm @login="onLogin" v-else/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
