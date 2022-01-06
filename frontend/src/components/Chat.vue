<template>
  <div class="card vh-100">
    <div class="card-header">
      <img alt="gRPC" src="@/assets/grpc.svg"> Web Chat
    </div>
    <div class="card-body overflow-scroll">
      <message-view
        v-for="(message, index) in messages"
        :key="index"
        :message="message"
        :class="index > 0 && 'mt-3'"
      />
      <p class="card-text" v-if="!messages.length">No Messages yet</p>
      <div ref="messagesBottom"></div>
    </div>
    <div class="card-footer">
      <submit-form-view @submit="onSubmit" v-if="isUserLoggedIn" />
      <login-form-view @submit="onLogin" v-else />
    </div>
  </div>
</template>

<script>
import Client from '@/Client'
import MessageView from './Message'
import LoginFormView from './LoginForm'
import SubmitFormView from './SubmitForm'

export default {
  data() {
    return {
      userName: null,
      messages: [],
    }
  },
  computed: {
    isUserLoggedIn() {
      return null !== this.userName
    }
  },
  mounted() {
    this.client = new Client()
    this.client.subscribe(this.onReceiveMessage)
  },
  methods: {
    addMessage(author, text, isOur) {
      this.messages.push({author, text, isOur})
      this.$nextTick(() => {
        // Code that will run only after the entire view has been re-rendered
        this.$refs.messagesBottom.scrollIntoView({behavior: 'smooth'})
      })
    },
    onReceiveMessage(author, text) {
      this.addMessage(author, text, false)
    },
    onSubmit(text) {
      this.addMessage(this.userName, text, true)
      this.client.sendMessage(this.userName, text)
    },
    onLogin(userName) {
      this.userName = userName
    }
  },
  components: {
    MessageView,
    LoginFormView,
    SubmitFormView
  }
}
</script>
