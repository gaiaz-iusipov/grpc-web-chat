package config

import (
	"context"
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/subosito/gotenv"
)

type Config struct {
	GRPCPort  uint16 `env:"GRPC_PORT"`
	HTTPPort  uint16 `env:"HTTP_PORT"`
	DebugPort uint16 `env:"DEBUG_PORT"`
}

func New() (*Config, error) {
	_ = gotenv.Load(".env")

	cfg := new(Config)
	if err := env.ParseWithOptions(cfg, env.Options{RequiredIfNoDef: true}); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return cfg, nil
}

func GRPCPort(ctx context.Context) uint16 {
	cfg, ok := ctx.Value(ctxKey{}).(*Config)
	if !ok {
		return 0
	}
	return cfg.GRPCPort
}

func HTTPPort(ctx context.Context) uint16 {
	cfg, ok := ctx.Value(ctxKey{}).(*Config)
	if !ok {
		return 0
	}
	return cfg.HTTPPort
}

func DebugPort(ctx context.Context) uint16 {
	cfg, ok := ctx.Value(ctxKey{}).(*Config)
	if !ok {
		return 0
	}
	return cfg.DebugPort
}
