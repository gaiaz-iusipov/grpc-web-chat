<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col col-md-6 offset-md-3 pt-3">
        <chat-view />
      </div>
    </div>
  </div>
</template>

<script>
import ChatView from '@/components/Chat'

export default {
  name: 'app',
  components: {
    ChatView
  }
}
</script>

<style src="bootstrap/dist/css/bootstrap.css"></style>
