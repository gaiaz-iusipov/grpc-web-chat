package grpcpublic

import (
	chatpb "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

type Controller interface {
	chatpb.ChatServer
}
