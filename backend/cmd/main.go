package main

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/grpc"
	"github.com/gaiaz-iusipov/grpc-web-chat/internal/http"
)

type config struct {
	GRPCPort  uint16 `env:"GRPC_PORT"`
	DebugPort uint16 `env:"DEBUG_PORT"`
}

func main() {
	_ = gotenv.Load()

	cfg := new(config)
	if err := env.Parse(cfg, env.Options{RequiredIfNoDef: true}); err != nil {
		log.Fatal().Err(err).Msg("env.Parse()")
	}

	grpcServer := grpc.New(cfg.GRPCPort)
	go func() {
		if err := grpcServer.Run(); err != nil {
			log.Fatal().Err(err).Msg("grpcServer.Run()")
		}
	}()

	debugServer := http.New(cfg.DebugPort)
	go func() {
		if err := debugServer.Run(); err != nil {
			log.Fatal().Err(err).Msg("debugServer.Run()")
		}
	}()

	log.Debug().Msg("app started")

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	log.Debug().Msg("shutting down gracefully")

	closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		wgErr := grpcServer.Close(closeCtx)
		if wgErr != nil {
			log.Error().Err(wgErr).Msg("grpcServer.Close()")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		wgErr := debugServer.Close(closeCtx)
		if wgErr != nil {
			log.Error().Err(wgErr).Msg("debugServer.Close()")
		}
	}()

	wg.Wait()
}
