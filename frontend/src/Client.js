import { v4 as uuidv4 } from 'uuid'
import { Subscribe, Client, Message, AddMessage } from './chat/v1/chat_pb'
import { ChatClient } from './chat/v1/chat_grpc_web_pb'

export default class {
  #protoClient;
  #clientUuid;
  constructor() {
    this.#protoClient = new ChatClient(process.env.VUE_APP_API_URL, null, null)
    this.#clientUuid = uuidv4()
  }
  subscribe(callback) {
    const req = new Subscribe.Request()
    req.setClientUuid(this.#clientUuid)

    const stream = this.#protoClient.subscribe(req)

    stream.on('data', resp => {
      const msg = resp.getMessage()
      callback(msg.getClient().getName(), msg.getText())
    })
  }
  sendMessage(author, text) {
    const client = new Client()
    client.setUuid(this.#clientUuid)
    client.setName(author)

    const msg = new Message()
    msg.setClient(client)
    msg.setText(text)

    const req = new AddMessage.Request()
    req.setMessage(msg)

    this.#protoClient.addMessage(req, {}, (err) => {
      if (null !== err) {
        throw err
      }
    })
  }
}
