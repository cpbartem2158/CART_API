package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed", s.logger)
		return
	}

	idStr := r.PathValue("id")
	cartId, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid cart ID", s.logger)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	cart, err := s.service.GetCart(ctx, cartId)
	if err != nil {
		if errors.Is(err, errorsx.ErrCartNotFound) {
			writeJSONError(w, http.StatusNotFound, "cart not found", s.logger)
			return
		}

		s.logger.Error("error view cart", err, s.logger)
		writeJSONError(w, http.StatusInternalServerError, "error view cart", s.logger)
		return
	}
	writeJSONResponse(w, http.StatusOK, cart, s.logger)
}
