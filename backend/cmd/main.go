package main

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/app"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(logHandler))

	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run app", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	initCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	appSvc, err := app.New(initCtx)
	if err != nil {
		return errors.Wrap(err, "init app")
	}

	appSvc.Run()

	slog.InfoContext(initCtx, "started")

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	slog.InfoContext(closeCtx, "shutting down")

	closeErr := appSvc.Close(closeCtx)
	if closeErr != nil {
		return errors.Wrap(closeErr, "close app")
	}

	return nil
}
