package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"

	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

func (s *Service) AddMessage(_ context.Context, message *proto.Message) (*emptypb.Empty, error) {
	log.Debug().
		Str("client_uuid", message.Client.Id).
		Str("message_text", message.Text).
		Msg("message received")

	for clientID, channel := range s.channels {
		if clientID == message.Client.Id {
			continue
		}

		channel <- message
	}

	return &emptypb.Empty{}, nil
}
