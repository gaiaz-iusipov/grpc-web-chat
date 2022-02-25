package service

import (
	"io"

	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

type Service struct {
	proto.UnimplementedChatServer
	closeCh  chan struct{}
	channels map[string]chan *proto.Message
}

var (
	_ proto.ChatServer = (*Service)(nil)
	_ io.Closer        = (*Service)(nil)
)

func New() *Service {
	return &Service{
		closeCh:  make(chan struct{}),
		channels: make(map[string]chan *proto.Message),
	}
}

func (s *Service) Close() error {
	close(s.closeCh)
	return nil
}
