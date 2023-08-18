package httpprivateserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	httpprivate "github.com/gaiaz-iusipov/grpc-web-chat/internal/http/private"
)

type Server struct {
	httpServer *http.Server
	running    chan struct{}
}

func New(
	listenPort uint16,
	controller httpprivate.Controller,
) Server {
	mux := chi.NewRouter()
	server := Server{
		httpServer: &http.Server{
			Addr:              ":" + strconv.FormatInt(int64(listenPort), 10),
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
		},
		running: make(chan struct{}),
	}

	// pprof
	mux.Mount("/debug", middleware.Profiler())

	// Prometheus
	mux.Method(http.MethodGet, "/metrics", promhttp.Handler())

	// Build Info
	mux.Get("/info", controller.GetInfo)

	// Kubernetes Probes
	mux.MethodFunc(http.MethodGet, "/live", controller.GetLive)
	mux.MethodFunc(http.MethodGet, "/ready", controller.GetReady)

	return server
}

func (s Server) Run() error {
	listener, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		return fmt.Errorf("net listen: %w", err)
	}

	close(s.running)

	err = s.httpServer.Serve(listener)
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http serve: %w", err)
	}
	return nil
}

func (s Server) Running() <-chan struct{} {
	return s.running
}

func (s Server) Close(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
