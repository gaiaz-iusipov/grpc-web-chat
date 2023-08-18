package appconfig

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/subosito/gotenv"
)

type Config struct {
	HTTPPublicPort  uint16 `env:"HTTP_PUBLIC_PORT"`
	GRPCPublicPort  uint16 `env:"GRPC_PUBLIC_PORT"`
	HTTPPrivatePort uint16 `env:"HTTP_PRIVATE_PORT"`
}

func New() (*Config, error) {
	_ = gotenv.Load(".env")

	cfg := new(Config)
	if err := env.ParseWithOptions(cfg, env.Options{RequiredIfNoDef: true}); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return cfg, nil
}
