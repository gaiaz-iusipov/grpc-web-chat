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
	"google.golang.org/grpc"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/debug"
	"github.com/gaiaz-iusipov/grpc-web-chat/internal/public"
)

type config struct {
	GRPCPort  uint16 `env:"GRPC_PORT"`
	HTTPPort  uint16 `env:"HTTP_PORT"`
	DebugPort uint16 `env:"DEBUG_PORT"`
}

func main() {
	_ = gotenv.Load()

	cfg := new(config)
	if err := env.Parse(cfg, env.Options{RequiredIfNoDef: true}); err != nil {
		log.Fatal().Err(err).Msg("env.Parse()")
	}

	grpcServer := grpc.NewServer()

	publicGRPCServer := public.NewGRPCServer(cfg.GRPCPort, grpcServer)
	go func() {
		if err := publicGRPCServer.Run(); err != nil {
			log.Fatal().Err(err).Msg("publicGRPCServer.Run()")
		}
	}()

	publicHTTPServer := public.NewHTTPServer(cfg.HTTPPort, grpcServer)
	go func() {
		if err := publicHTTPServer.Run(); err != nil {
			log.Fatal().Err(err).Msg("publicHTTPServer.Run()")
		}
	}()

	debugHTTPServer := debug.NewHTTPServer(cfg.DebugPort)
	go func() {
		if err := debugHTTPServer.Run(); err != nil {
			log.Fatal().Err(err).Msg("debugHTTPServer.Run()")
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

		wgErr := publicGRPCServer.Close(closeCtx)
		if wgErr != nil {
			log.Error().Err(wgErr).Msg("publicGRPCServer.Close()")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		wgErr := publicHTTPServer.Close(closeCtx)
		if wgErr != nil {
			log.Error().Err(wgErr).Msg("publicHTTPServer.Close()")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		wgErr := debugHTTPServer.Close(closeCtx)
		if wgErr != nil {
			log.Error().Err(wgErr).Msg("debugHTTPServer.Close()")
		}
	}()

	wg.Wait()
}
