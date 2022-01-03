package grpc

import (
	"context"
	"net"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/grpc/service"
	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

type Server struct {
	port       uint16
	grpcServer *grpc.Server
}

func New(port uint16) *Server {
	server := grpc.NewServer()
	proto.RegisterChatServer(server, service.New())

	return &Server{
		port:       port,
		grpcServer: server,
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(s.port), 10))
	if err != nil {
		return errors.Wrap(err, "net.Listen()")
	}

	err = s.grpcServer.Serve(listener)
	if err != nil {
		return errors.Wrap(err, "server.Serve()")
	}
	return nil
}

func (s *Server) Close(ctx context.Context) error {
	stopped := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "grpcServer.GracefulStop()")
	case <-stopped:
	}
	return nil
}
