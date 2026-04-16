package handlers

import (
	"log/slog"
	"net/http"

	"github.com/cpbartem2158/CART_API/internal/config"
	"github.com/cpbartem2158/CART_API/internal/service"
)

type Server struct {
	service service.Servicer
	logger  *slog.Logger
	config  *config.ServerConfig
}

func NewServer(service service.Servicer, logger *slog.Logger, config *config.ServerConfig) *Server {
	return &Server{service: service, logger: logger, config: config}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /carts", s.CreateCart)
	mux.HandleFunc("POST /carts/{id}/items", s.AddCartItemToCart)
	mux.HandleFunc("DELETE /carts/{id}/items/{item_id}", s.RemoveCartItem)

	httpServer := &http.Server{
		Addr:         ":" + s.config.Port,
		Handler:      mux,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
		IdleTimeout:  s.config.IdleTimeout,
	}
	s.logger.Info("starting HTTP server on port " + s.config.Port)
	return httpServer.ListenAndServe()
}
