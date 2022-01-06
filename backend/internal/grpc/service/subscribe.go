package service

import (
	"github.com/rs/zerolog/log"

	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

func (s *Service) Subscribe(req *proto.SubscribeRequest, respSender proto.Chat_SubscribeServer) error {
	clientUUID := req.GetClientUuid()

	channel := make(chan *proto.Message)
	s.channels[clientUUID] = channel

	log.Debug().
		Str("client_uuid", clientUUID).
		Msg("client subscribed")

	for message := range channel {
		err := respSender.Send(&proto.SubscribeResponse{
			Message: message,
		})
		if err != nil {
			delete(s.channels, clientUUID)
			log.Debug().
				Str("client_uuid", clientUUID).
				Msg("client unsubscribed")
			break
		}
	}

	return nil
}
