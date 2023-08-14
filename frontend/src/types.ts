export interface IUser {
  id: string
  name: string
}

export interface IMessage {
  author: IUser
  text: string
  isOur: boolean
  createdAt: Date
}
