package config

import "context"

type ctxKey struct{}

func (c *Config) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, c)
}
