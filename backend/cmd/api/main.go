// Command api starts the who-has-my-book HTTP API.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sidnei-junior/who-has-my-book/backend/internal/platform/config"
	"github.com/sidnei-junior/who-has-my-book/backend/internal/platform/database"
	"github.com/sidnei-junior/who-has-my-book/backend/internal/platform/logger"
	"github.com/sidnei-junior/who-has-my-book/backend/internal/platform/server"
)

func main() {
	if err := run(); err != nil {
		slog.Error("application error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	log := logger.New(cfg.LogLevel)
	slog.SetDefault(log)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := database.Open(ctx, cfg.DatabaseURL, cfg.IsDevelopment() && cfg.LogLevel == "debug")
	if err != nil {
		return err
	}
	defer database.Close(db)
	log.Info("database connected")

	srv := server.New(cfg.HTTPPort, cfg.CORSAllowedOrigins, db, log)

	errCh := make(chan error, 1)
	go func() { errCh <- srv.Start() }()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Info("shutting down")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}
