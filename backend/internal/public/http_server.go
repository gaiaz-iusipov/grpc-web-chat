package public

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/app/config"
)

type HTTPServer struct {
	port   uint16
	server *http.Server
}

func NewHTTPServer(ctx context.Context, grpcServer *grpc.Server) (HTTPServer, error) {
	port := config.HTTPPort(ctx)
	if port == 0 {
		return HTTPServer{}, errors.New("missing HTTPPort")
	}

	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(_ string) bool {
		return true
	}))

	return HTTPServer{
		port: port,
		server: &http.Server{
			Handler: wrappedGrpc,
		},
	}, nil
}

func (s HTTPServer) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	err = s.server.Serve(ln)
	if err != http.ErrServerClosed {
		return errors.Wrap(err, "server.Serve()")
	}
	return nil
}

func (s HTTPServer) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	return errors.Wrap(err, "server.Shutdown()")
}
