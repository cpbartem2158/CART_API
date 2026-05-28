package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/cpbartem2158/CART_API/internal/config"
	"github.com/cpbartem2158/CART_API/internal/service"
)

type Server struct {
	service    service.Servicer
	logger     *slog.Logger
	config     *config.ServerConfig
	httpServer *http.Server
}

func NewServer(service service.Servicer, logger *slog.Logger, config *config.ServerConfig) *Server {
	mux := http.NewServeMux()

	s := &Server{
		service: service,
		logger:  logger,
		config:  config,
	}

	mux.HandleFunc("POST /carts", s.CreateCart)
	mux.HandleFunc("POST /carts/{id}/items", s.AddCartItemToCart)
	mux.HandleFunc("DELETE /carts/{id}/items/{item_id}", s.RemoveCartItem)
	mux.HandleFunc("GET /carts/{id}", s.GetCart)
	mux.HandleFunc("GET /carts/{id}/price", s.CalculatePrice)

	if strings.TrimSpace(config.Port) == "" {
		config.Port = "8080"
	}
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      mux,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	return s
}

func (s *Server) Start() error {
	s.logger.Info("http server listening", "addr", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down HTTP server")
	return s.httpServer.Shutdown(ctx)
}
