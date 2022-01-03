package main

import (
	"github.com/rs/zerolog/log"
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
	log.Debug().Str("client_uuid", client.Id).Msg("client subscribed")

	for message := range channel {
		if err := cs.Send(message); err != nil {
			delete(server.channels, client.Id)
			log.Debug().Str("client_uuid", client.Id).Msg("client unsubscribed")
			break
		}
	}

	return nil
}

func (server *Server) AddMessage(_ context.Context, message *proto.Message) (*emptypb.Empty, error) {
	log.Debug().Str("message_text", message.Text).Msg("message received")

	for clientID, channel := range server.channels {
		if clientID == message.Client.Id {
			continue
		}

		channel <- message
	}

	return &emptypb.Empty{}, nil
}
