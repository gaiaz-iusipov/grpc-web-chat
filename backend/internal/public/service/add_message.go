package service

import (
	"context"

	"github.com/rs/zerolog/log"

	chatv1 "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

func (s *Service) AddMessage(ctx context.Context, req *chatv1.AddMessage_Request) (*chatv1.AddMessage_Response, error) {
	message := req.GetMessage()

	log.Ctx(ctx).Debug().
		Str("client_uuid", message.Client.Uuid).
		Str("message_text", message.Text).
		Msg("message received")

	for clientUUID, channel := range s.channels {
		if clientUUID == message.Client.Uuid {
			continue
		}

		channel <- message
	}

	return &chatv1.AddMessage_Response{}, nil
}
