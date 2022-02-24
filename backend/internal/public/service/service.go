package service

import (
	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

type Service struct {
	proto.UnimplementedChatServer
	channels map[string]chan *proto.Message
}

var _ proto.ChatServer = (*Service)(nil)

func New() *Service {
	return &Service{
		channels: make(map[string]chan *proto.Message),
	}
}
