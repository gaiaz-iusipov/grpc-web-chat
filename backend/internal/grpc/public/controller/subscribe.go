package grpcpubliccontroller

import (
	"fmt"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	chatv1 "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

func (c *Controller) Subscribe(req *chatv1.Subscribe_Request, respSender chatv1.Chat_SubscribeServer) (err error) {
	ctx := respSender.Context()

	clientUUID := req.GetClientUuid()
	channel := make(chan *chatv1.Message)

	c.channelsMu.Lock()
	c.channels[clientUUID] = channel
	c.channelsMu.Unlock()

	slog.InfoContext(ctx, "client subscribed", "client_uuid", clientUUID)

	defer func() {
		c.channelsMu.Lock()
		delete(c.channels, clientUUID)
		c.channelsMu.Unlock()

		slog.InfoContext(ctx, "client unsubscribed", "client_uuid", clientUUID)
	}()

	for {
		select {
		case msg := <-channel:
			err = respSender.Send(&chatv1.Subscribe_Response{
				Message: msg,
			})
			if err != nil {
				if status.Code(err) == codes.Unavailable {
					return nil
				}

				return fmt.Errorf("send message: %w", err)
			}
		case <-ctx.Done():
			return
		case <-c.closeCh:
			return
		}
	}
}
