package main

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/app"
)

func main() {
	initCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	appSvc, err := app.New(initCtx)
	if err != nil {
		log.Fatal().Err(err).Msg("app.New()")
	}

	appSvc.Run()

	log.Debug().Msg("app started")

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	log.Debug().Msg("shutting down gracefully")

	closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = appSvc.Close(closeCtx)
	if err != nil {
		log.Error().Err(err).Msg("appSvc.Close()")
	}
}
