package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-deeper/app"

	appconfig "github.com/gaiaz-iusipov/grpc-web-chat/internal/app/config"
	grpcpubliccontroller "github.com/gaiaz-iusipov/grpc-web-chat/internal/grpc/public/controller"
	grpcpublicserver "github.com/gaiaz-iusipov/grpc-web-chat/internal/grpc/public/server"
	httpprivatecontroller "github.com/gaiaz-iusipov/grpc-web-chat/internal/http/private/controller"
	httpprivateserver "github.com/gaiaz-iusipov/grpc-web-chat/internal/http/private/server"
	httppublicserver "github.com/gaiaz-iusipov/grpc-web-chat/internal/http/public/server"
)

func Run(ctx context.Context) error {
	initCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	config, err := appconfig.New()
	if err != nil {
		return fmt.Errorf("new config: %w", err)
	}

	slog.LogAttrs(ctx, slog.LevelInfo, "starting", app.LogAttrs()...)

	httpPrivateController := httpprivatecontroller.New()

	httpPrivateServer := httpprivateserver.New(
		config.HTTPPrivatePort,
		httpPrivateController,
	)

	if err = runAsync(initCtx, httpPrivateServer.Run, httpPrivateServer.Running); err != nil {
		return fmt.Errorf("run http private server: %w", err)
	}

	grpcPublicController := grpcpubliccontroller.New()

	grpcPublicServer := grpcpublicserver.New(config.GRPCPublicPort, grpcPublicController)

	httpPublicServer := httppublicserver.New(
		config.HTTPPublicPort,
		grpcPublicServer,
	)

	if err = runAsync(initCtx, httpPublicServer.Run, httpPublicServer.Running); err != nil {
		return fmt.Errorf("run http public server: %w", err)
	}

	httpPrivateController.SetReady(true)

	slog.InfoContext(ctx, "started",
		"http_public_port", config.HTTPPublicPort,
		"grpc_public_port", config.GRPCPublicPort,
		"http_private_port", config.HTTPPrivatePort,
	)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	httpPrivateController.SetReady(false)

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	slog.InfoContext(closeCtx, "shutting down")

	grpcPublicController.Close()

	if closeErr := httpPublicServer.Close(closeCtx); closeErr != nil {
		slog.ErrorContext(closeCtx, "failed to close http public server", "error", closeErr)
	}

	if closeErr := grpcPublicServer.Close(closeCtx); closeErr != nil {
		slog.ErrorContext(closeCtx, "failed to close grpc public server", "error", closeErr)
	}

	if closeErr := httpPrivateServer.Close(closeCtx); closeErr != nil {
		slog.ErrorContext(closeCtx, "failed to close http private server", "error", closeErr)
	}

	return nil
}

func runAsync(ctx context.Context, runFn func() error, okFn func() <-chan struct{}) error {
	errCh := make(chan error)
	go func() {
		errCh <- runFn()
	}()

	var err error
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-errCh:
	case <-okFn():
	}
	return err
}
