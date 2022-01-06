package service

import (
	"context"

	"github.com/rs/zerolog/log"

	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

func (s *Service) AddMessage(_ context.Context, req *proto.AddMessageRequest) (*proto.AddMessageResponse, error) {
	message := req.GetMessage()

	log.Debug().
		Str("client_uuid", message.Client.Uuid).
		Str("message_text", message.Text).
		Msg("message received")

	for clientUUID, channel := range s.channels {
		if clientUUID == message.Client.Uuid {
			continue
		}

		channel <- message
	}

	return &proto.AddMessageResponse{}, nil
}
