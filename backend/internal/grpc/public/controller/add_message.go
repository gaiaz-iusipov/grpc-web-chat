package grpcpubliccontroller

import (
	"context"
	"log/slog"

	chatv1 "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

func (c *Controller) AddMessage(ctx context.Context, req *chatv1.AddMessage_Request) (*chatv1.AddMessage_Response, error) {
	message := req.GetMessage()

	slog.InfoContext(ctx, "message received",
		"client_uuid", message.Client.Uuid,
		"message_text", message.Text,
	)

	c.channelsMu.RLock()
	defer c.channelsMu.RUnlock()

	for clientUUID, channel := range c.channels {
		if clientUUID == message.Client.Uuid {
			continue
		}

		channel <- message
	}

	return &chatv1.AddMessage_Response{}, nil
}
