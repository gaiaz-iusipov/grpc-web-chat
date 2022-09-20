package main

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/app"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Ctx(ctx).Fatal().Err(err)
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

	log.Ctx(initCtx).Debug().Msg("app started")

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Ctx(closeCtx).Debug().Msg("shutting down gracefully")

	closeErr := appSvc.Close(closeCtx)
	if closeErr != nil {
		return errors.Wrap(closeErr, "close app")
	}

	return nil
}
