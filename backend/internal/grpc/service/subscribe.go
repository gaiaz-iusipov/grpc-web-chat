package service

import (
	"github.com/rs/zerolog/log"

	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

func (s *Service) Subscribe(client *proto.Client, cs proto.Chat_SubscribeServer) error {
	channel := make(chan *proto.Message)
	s.channels[client.Id] = channel
	log.Debug().Str("client_uuid", client.Id).Msg("client subscribed")

	for message := range channel {
		if err := cs.Send(message); err != nil {
			delete(s.channels, client.Id)
			log.Debug().Str("client_uuid", client.Id).Msg("client unsubscribed")
			break
		}
	}

	return nil
}
