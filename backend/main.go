package main

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	proto "github.com/gaiaz-iusipov/grpc-web-chat/chat"
)

const (
	port = 3000
)

func main() {
	log.SetLevel(log.DebugLevel)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() {
		if err := lis.Close(); err != nil {
			log.Fatalf("failed to close: %v", err)
		}
	}()

	server := grpc.NewServer()
	proto.RegisterChatServer(server, NewServer())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
