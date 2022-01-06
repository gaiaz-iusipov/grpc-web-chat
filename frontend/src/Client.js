import { v4 as uuidv4 } from 'uuid'
import {Client, Message} from './chat/chat_pb'
import {ChatClient} from './chat/chat_grpc_web_pb'

export default class {
  #client;
  #protoClient;
  constructor() {
    this.#client = new Client()
    this.#client.setId(uuidv4())
    this.#protoClient = new ChatClient(process.env.VUE_APP_API_URL, null, null)
  }
  subscribe(callback) {
    const stream = this.#protoClient.subscribe(this.#client)

    stream.on('data', response => {
      callback(response.getAuthor(), response.getText())
    })
  }
  sendMessage(author, text) {
    const request = new Message()
    request.setClient(this.#client)
    request.setAuthor(author)
    request.setText(text)

    this.#protoClient.addMessage(request, {}, (err) => {
      if (null !== err) {
        throw err
      }
    })
  }
}
