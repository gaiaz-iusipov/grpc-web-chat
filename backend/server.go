package main

import (
	proto "github.com/gaiaz-iusipov/grpc-web-chat/chat"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Server struct {
	channels map[string]chan *proto.Message
}

func NewServer() *Server {
	server := new(Server)
	server.channels = make(map[string]chan *proto.Message)

	return server
}

func (s *Server) Subscribe(client *proto.Client, cs proto.Chat_SubscribeServer) error {
	channel := make(chan *proto.Message)
	s.channels[client.Id] = channel
	log.Debugf("added a new channel: %s", client.Id)

	for message := range channel {
		log.Debug("retrieved a message from a channel")

		if err := cs.Send(message); err != nil {
			delete(s.channels, client.Id)
			log.Debugf("removed a channel: %s", client.Id)

			break
		}
	}

	return nil
}

func (s *Server) AddMessage(ctx context.Context, message *proto.Message) (*empty.Empty, error) {
	log.Debugf("received a new message: %s", message.Text)

	for clientId, channel := range s.channels {
		if clientId == message.Client.Id {
			continue
		}

		channel <- message
	}

	return &empty.Empty{}, nil
}
