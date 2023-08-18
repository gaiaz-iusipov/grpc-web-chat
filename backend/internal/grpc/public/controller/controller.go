package grpcpubliccontroller

import (
	"sync"

	grpcpublic "github.com/gaiaz-iusipov/grpc-web-chat/internal/grpc/public"
	chatpb "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

type Controller struct {
	chatpb.UnimplementedChatServer
	channels   map[string]chan *chatpb.Message
	channelsMu sync.RWMutex
	closeCh    chan struct{}
}

var _ grpcpublic.Controller = (*Controller)(nil)

func New() *Controller {
	return &Controller{
		channels: make(map[string]chan *chatpb.Message),
		closeCh:  make(chan struct{}),
	}
}

func (c *Controller) Close() {
	close(c.closeCh)
}
