package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
)

type Server struct {
	httpServer *http.Server
}

func New(port uint16) *Server {
	server := &Server{
		httpServer: &http.Server{
			Addr: ":" + strconv.FormatInt(int64(port), 10),
		},
	}

	router := chi.NewMux()
	router.Mount("/debug", middleware.Profiler())
	router.Get("/ready", server.ready)
	router.Get("/live", server.live)
	server.httpServer.Handler = router
	return server
}

func (s *Server) Run() error {
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return errors.Wrap(err, "httpServer.ListenAndServe()")
	}
	return nil
}

func (s *Server) Close(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	return errors.Wrap(err, "httpServer.Shutdown()")
}

func (Server) ready(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("{}"))
}

func (Server) live(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("{}"))
}
