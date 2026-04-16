package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (s *Server) CalculatePrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed", s.logger)
		return
	}
	cartIdStr := r.PathValue("id")
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid cart ID", s.logger)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cartPrice, err := s.service.CalculatePrice(ctx, cartId)
	if err != nil {
		if errors.Is(err, errorsx.ErrCartNotFound) {
			writeJSONError(w, http.StatusNotFound, "cart not found", s.logger)
			return
		}
		s.logger.Error("failed to calculate price", err)
		writeJSONError(w, http.StatusInternalServerError, "failed to calculate price", s.logger)
		return
	}
	writeJSONResponse(w, http.StatusOK, cartPrice, s.logger)

}
