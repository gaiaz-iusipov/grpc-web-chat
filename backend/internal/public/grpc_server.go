package public

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/app/config"
	"github.com/gaiaz-iusipov/grpc-web-chat/internal/public/service"
	chatv1 "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

type GRPCServer struct {
	port   uint16
	server *grpc.Server
}

func NewGRPCServer(ctx context.Context, grpcServer *grpc.Server, chatSrv *service.Service) (GRPCServer, error) {
	port := config.GRPCPort(ctx)
	if port == 0 {
		return GRPCServer{}, errors.New("missing GRPCPort")
	}

	chatv1.RegisterChatServer(grpcServer, chatSrv)

	return GRPCServer{
		port:   port,
		server: grpcServer,
	}, nil
}

func (s GRPCServer) Run() error {
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

func (s GRPCServer) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	stopCh := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(stopCh)
	}()

	select {
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "server.GracefulStop()")
	case <-stopCh:
	}
	return nil
}
