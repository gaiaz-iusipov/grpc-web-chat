package main

import (
	"context"
	"log/slog"
	_ "net/http/pprof"
	"os"

	"github.com/gaiaz-iusipov/grpc-web-chat/internal/app"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(logHandler))

	ctx := context.Background()
	if err := app.Run(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run app", "error", err)
		os.Exit(1)
	}
}
