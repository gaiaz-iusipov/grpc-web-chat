package public

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type HTTPServer struct {
	port   uint16
	server *http.Server
}

func NewHTTPServer(port uint16, grpcServer *grpc.Server) *HTTPServer {
	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(_ string) bool {
		return true
	}))

	return &HTTPServer{
		port: port,
		server: &http.Server{
			Handler: wrappedGrpc,
		},
	}
}

func (s *HTTPServer) Run() error {
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

func (s *HTTPServer) Close(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	return errors.Wrap(err, "server.Shutdown()")
}
