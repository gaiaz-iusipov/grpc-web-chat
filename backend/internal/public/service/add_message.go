package service

import (
	"context"

	"golang.org/x/exp/slog"

	chatv1 "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

func (s *Service) AddMessage(ctx context.Context, req *chatv1.AddMessage_Request) (*chatv1.AddMessage_Response, error) {
	message := req.GetMessage()

	slog.InfoContext(ctx, "message received",
		"client_uuid", message.Client.Uuid,
		"message_text", message.Text,
	)

	s.channelsMu.RLock()
	defer s.channelsMu.RUnlock()

	for clientUUID, channel := range s.channels {
		if clientUUID == message.Client.Uuid {
			continue
		}

		channel <- message
	}

	return &chatv1.AddMessage_Response{}, nil
}
