package app

import (
	"context"
	"io"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/app/config"
	"github.com/gaiaz-iusipov/grpc-web-chat/internal/debug"
	"github.com/gaiaz-iusipov/grpc-web-chat/internal/public"
	"github.com/gaiaz-iusipov/grpc-web-chat/internal/public/service"
	chatv1 "github.com/gaiaz-iusipov/grpc-web-chat/pkg/chat/v1"
)

type ClosableChatServer interface {
	chatv1.ChatServer
	io.Closer
}

type App struct {
	chatSrv          ClosableChatServer
	publicGRPCServer public.GRPCServer
	publicHTTPServer public.HTTPServer
	debugHTTPServer  debug.HTTPServer
}

func New(ctx context.Context) (App, error) {
	cfg, err := config.New()
	if err != nil {
		return App{}, errors.Wrap(err, "config.New()")
	}

	ctx = cfg.ToContext(ctx)

	grpcServer := grpc.NewServer()
	chatSrv := service.New()

	publicGRPCServer, err := public.NewGRPCServer(ctx, grpcServer, chatSrv)
	if err != nil {
		return App{}, errors.Wrap(err, "public.NewGRPCServer()")
	}

	publicHTTPServer, err := public.NewHTTPServer(ctx, grpcServer)
	if err != nil {
		return App{}, errors.Wrap(err, "public.NewHTTPServer()")
	}

	debugHTTPServer, err := debug.NewHTTPServer(ctx)
	if err != nil {
		return App{}, errors.Wrap(err, "debug.NewHTTPServer()")
	}

	return App{
		chatSrv:          chatSrv,
		publicGRPCServer: publicGRPCServer,
		publicHTTPServer: publicHTTPServer,
		debugHTTPServer:  debugHTTPServer,
	}, nil
}

func (a App) Run() {
	go func() {
		if err := a.publicGRPCServer.Run(); err != nil {
			log.Fatal().Err(err).Msg("publicGRPCServer.Run()")
		}
	}()

	go func() {
		if err := a.publicHTTPServer.Run(); err != nil {
			log.Fatal().Err(err).Msg("publicHTTPServer.Run()")
		}
	}()

	go func() {
		if err := a.debugHTTPServer.Run(); err != nil {
			log.Fatal().Err(err).Msg("debugHTTPServer.Run()")
		}
	}()
}

func (a App) Close(ctx context.Context) error {
	_ = a.chatSrv.Close()

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		egErr := a.publicHTTPServer.Close(egCtx)
		if egErr != nil {
			return errors.Wrap(egErr, "publicHTTPServer.Close()")
		}

		egErr = a.publicGRPCServer.Close(egCtx)
		return errors.Wrap(egErr, "publicGRPCServer.Close()")
	})

	eg.Go(func() error {
		egErr := a.debugHTTPServer.Close(egCtx)
		return errors.Wrap(egErr, "debugHTTPServer.Close()")
	})

	return eg.Wait()
}
