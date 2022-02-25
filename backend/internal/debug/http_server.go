package debug

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/app/config"
)

type HTTPServer struct {
	port   uint16
	server *http.Server
}

func NewHTTPServer(ctx context.Context) (HTTPServer, error) {
	port := config.DebugPort(ctx)
	if port == 0 {
		return HTTPServer{}, errors.New("missing DebugPort")
	}

	s := HTTPServer{
		port:   port,
		server: &http.Server{},
	}

	router := chi.NewMux()
	router.Mount("/debug", middleware.Profiler())
	router.Get("/ready", s.ready)
	router.Get("/live", s.live)
	s.server.Handler = router
	return s, nil
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

func (HTTPServer) ready(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("{}"))
}

func (HTTPServer) live(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("{}"))
}
