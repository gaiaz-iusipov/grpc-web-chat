<template>
  <div class="container">
    <div class="row">
      <div class="col col-md-6 offset-md-3">
        <chat-view />
      </div>
    </div>
  </div>
</template>

<script>
import ChatView from './components/Chat.vue'

export default {
  name: 'App',
  components: {
    ChatView
  }
}
</script>

<style src="bootstrap/dist/css/bootstrap.css"></style>
