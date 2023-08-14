import type { IChatClient } from './chat/v1/chat.client'
import { ChatClient } from './chat/v1/chat.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { AddMessage_Request, Subscribe_Request } from './chat/v1/chat'
import type { IUser } from '@/types'

export default class {
  private readonly grpcClient: IChatClient

  constructor(apiUrl: string) {
    const transport = new GrpcWebFetchTransport({
      baseUrl: apiUrl
    })
    this.grpcClient = new ChatClient(transport)
  }

  async subscribe(userUuid: string, callback: any): Promise<void> {
    const req: Subscribe_Request = {
      clientUuid: userUuid,
    }
    const stream = this.grpcClient.subscribe(req)

    for await (const resp of stream.responses) {
      if (!resp.message?.client) {
        continue
      }

      callback(resp.message)
    }
  }

  async sendMessage(user: IUser, text: string): Promise<void> {
    const req: AddMessage_Request = {
      message: {
        client: {
          uuid: user.id,
          name: user.name
        },
        text: text,
      }
    }

    try {
      await this.grpcClient.addMessage(req)
    } catch (e) {
      console.log(e)
    }
  }
}
