package handlers

import (
	"context"
	"net/http"
	"time"
)

func (s *Server) CreateCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed", s.logger)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	cart, err := s.service.CreateCart(ctx)
	if err != nil {
		s.logger.Error("failed to create cart", err)
		writeJSONError(w, http.StatusInternalServerError, "something went wrong", s.logger)
		return
	}

	writeJSONResponse(w, http.StatusOK, cart, s.logger)
}
