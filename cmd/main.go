package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cpbartem2158/CART_API/internal/config"
	"github.com/cpbartem2158/CART_API/internal/db"
	"github.com/cpbartem2158/CART_API/internal/handlers"
	"github.com/cpbartem2158/CART_API/internal/repository"
	"github.com/cpbartem2158/CART_API/internal/service"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.LoadConfig("config")
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, err := db.Connect(ctx, cfg.Database)
	if err != nil {
		logger.Error("failed to connect database", "error", err)
		os.Exit(1)
	}
	defer database.Close()

	repo := repository.NewRepository(database)
	service := service.NewService(repo, logger)
	server := handlers.NewServer(service, logger, &cfg.Server)

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- server.Start()
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		logger.Error("server error", "error", err)
		os.Exit(1)
	case sig := <-shutdown:
		logger.Info("Shutting down...", "signal", sig.String())

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("failed to shutdown", "error", err)
			os.Exit(1)
		}
		logger.Info("Shutdown complete")
	}

}
