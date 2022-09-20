package service

import (
	"io"

	chatv1 "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

type Service struct {
	chatv1.UnimplementedChatServer
	closeCh  chan struct{}
	channels map[string]chan *chatv1.Message
}

var (
	_ chatv1.ChatServer = (*Service)(nil)
	_ io.Closer         = (*Service)(nil)
)

func New() *Service {
	return &Service{
		closeCh:  make(chan struct{}),
		channels: make(map[string]chan *chatv1.Message),
	}
}

func (s *Service) Close() error {
	close(s.closeCh)
	return nil
}
