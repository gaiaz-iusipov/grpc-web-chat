package httppublicserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	httpServer *http.Server
	running    chan struct{}
}

func New(
	listenPort uint16,
	handler http.Handler,
) Server {
	server := Server{
		httpServer: &http.Server{
			Addr:              ":" + strconv.FormatInt(int64(listenPort), 10),
			Handler:           handler,
			ReadHeaderTimeout: 5 * time.Second,
		},
		running: make(chan struct{}),
	}

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
