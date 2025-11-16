package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/Marie20767/go-web-app-template/internal/config"
	"github.com/Marie20767/go-web-app-template/internal/server"
	"github.com/Marie20767/go-web-app-template/internal/store"
)

func main() {
	if err := run(); err != nil {
		slog.Error("run failed", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("shutting down gracefully...")
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.ParseEnv()
	if err != nil {
		return err
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}))
	slog.SetDefault(logger)

	db, err := store.NewStore(ctx, cfg.DbURL)
	if err != nil {
		return err
	}
	defer db.Close()
	slog.Info("connected to DB successfully")

	svr := server.New(db, cfg.Port)
	serverErr := make(chan error, 1)

	go func() {
		serverErr <- svr.Start()
	}()

	select {
	case <-ctx.Done(): // blocks until server error OR signal received (e.g. by ctrl-C or process killed)
		slog.Info("shutdown signal received")
	case svrErr := <-serverErr:
		err = svrErr
	}

	shutdownErr := svr.Stop()
	if shutdownErr != nil {
		slog.Error("server shutdown error", slog.Any("error", shutdownErr))
	}

	return err
}
