package public

import (
	"context"
	"fmt"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/public/service"
	proto "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat"
)

type GRPCServer struct {
	port   uint16
	server *grpc.Server
}

func NewGRPCServer(port uint16, grpcServer *grpc.Server) *GRPCServer {
	proto.RegisterChatServer(grpcServer, service.New())

	return &GRPCServer{
		port:   port,
		server: grpcServer,
	}
}

func (s *GRPCServer) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return errors.Wrap(err, "net.Listen()")
	}

	err = s.server.Serve(listener)
	if err != grpc.ErrServerStopped {
		return errors.Wrap(err, "server.Serve()")
	}
	return nil
}

func (s *GRPCServer) Close(ctx context.Context) error {
	stopped := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "server.GracefulStop()")
	case <-stopped:
	}
	return nil
}
