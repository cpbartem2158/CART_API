package main

import (
	"context"
	"log/slog"
	"os"
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
		logger.Error("failed to load config", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, err := db.Connect(ctx, cfg.Database)
	if err != nil {
		logger.Error("failed to connect database", err)
		return
	}
	defer database.Close()

	repo := repository.NewRepository(database)
	service := service.NewService(repo, logger)
	server := handlers.NewServer(service, logger, &cfg.Server)

	if err := server.Start(); err != nil {
		logger.Error("failed to start server", err)
		return
	}

}
