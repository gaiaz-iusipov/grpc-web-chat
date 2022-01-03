package main

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

const (
	port = 3000
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("net.Listen()")
	}
	defer lis.Close()

	server := grpc.NewServer()
	proto.RegisterChatServer(server, NewServer())

	err = server.Serve(lis)
	if err != nil {
		log.Fatal().Err(err).Msg("server.Serve()")
	}
}
