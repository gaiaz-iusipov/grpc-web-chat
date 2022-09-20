package service

import (
	"github.com/rs/zerolog/log"

	chatv1 "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

func (s *Service) Subscribe(req *chatv1.Subscribe_Request, respSender chatv1.Chat_SubscribeServer) error {
	clientUUID := req.GetClientUuid()

	channel := make(chan *chatv1.Message)
	s.channels[clientUUID] = channel

	log.Debug().
		Str("client_uuid", clientUUID).
		Msg("client subscribed")

	for {
		select {
		case msg := <-channel:
			err := respSender.Send(&chatv1.Subscribe_Response{
				Message: msg,
			})
			if err != nil {
				delete(s.channels, clientUUID)
				log.Ctx(respSender.Context()).Debug().
					Str("client_uuid", clientUUID).
					Msg("client unsubscribed")
				return nil
			}
		case <-s.closeCh:
			return nil
		}
	}
}
