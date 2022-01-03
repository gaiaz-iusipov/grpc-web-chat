package main

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"

	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

type Server struct {
	proto.UnimplementedChatServer
	channels map[string]chan *proto.Message
}

var _ proto.ChatServer = (*Server)(nil)

func NewServer() *Server {
	return &Server{
		channels: make(map[string]chan *proto.Message),
	}
}

func (server *Server) Subscribe(client *proto.Client, cs proto.Chat_SubscribeServer) error {
	channel := make(chan *proto.Message)
	server.channels[client.Id] = channel
	log.Debugf("added a new channel: %s", client.Id)

	for message := range channel {
		if err := cs.Send(message); err != nil {
			delete(server.channels, client.Id)
			log.Debugf("removed a channel: %s", client.Id)

			break
		}
	}

	return nil
}

func (server *Server) AddMessage(_ context.Context, message *proto.Message) (*emptypb.Empty, error) {
	log.Debugf("received a new message: %s", message.Text)

	for clientID, channel := range server.channels {
		if clientID == message.Client.Id {
			continue
		}

		channel <- message
	}

	return &emptypb.Empty{}, nil
}
