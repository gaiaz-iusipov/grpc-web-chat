package grpcpublicserver

import (
	"context"
	"net"
	"net/http"
	"strconv"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcpublic "github.com/gaiaz-iusipov/grpc-web-chat/internal/grpc/public"
	chatpb "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

type GRPCServer interface {
	Serve(net.Listener) error
	GracefulStop()
}

type Server struct {
	addr           string
	grpcServer     GRPCServer
	grpcWebHandler http.Handler
	running        chan struct{}
}

func New(listenPort uint16, controller grpcpublic.Controller) Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcprometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(
			grpcprometheus.StreamServerInterceptor,
		),
	)

	reflection.Register(grpcServer)
	chatpb.RegisterChatServer(grpcServer, controller)

	return Server{
		addr:       ":" + strconv.FormatInt(int64(listenPort), 10),
		grpcServer: grpcServer,
		grpcWebHandler: grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(_ string) bool {
			return true
		})),
		running: make(chan struct{}),
	}
}

func (s Server) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return errors.Wrap(err, "listen addr")
	}

	close(s.running)

	err = s.grpcServer.Serve(listener)
	if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return errors.Wrap(err, "grpc serve")
	}
	return nil
}

func (s Server) Running() <-chan struct{} {
	return s.running
}

func (s Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.grpcWebHandler.ServeHTTP(rw, req)
}

func (s Server) Close(ctx context.Context) error {
	stopCh := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(stopCh)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-stopCh:
	}
	return nil
}
