package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (s *Server) RemoveCartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed", s.logger)
		return
	}

	cartIdStr := r.PathValue("id")
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid cart ID", s.logger)
		return
	}
	itemIdStr := r.PathValue("item_id")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid item ID", s.logger)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()
	err = s.service.RemoveItem(ctx, cartId, itemId)
	if err != nil {
		if errors.Is(err, errorsx.ErrCartNotFound) {
			writeJSONError(w, http.StatusNotFound, "cart not found", s.logger)
			return
		}
		if errors.Is(err, errorsx.ErrCartItemNotFound) {
			writeJSONError(w, http.StatusNotFound, "cart item not found", s.logger)
			return
		}
		s.logger.Error("failed to remove cart item", "error", err)
		writeJSONError(w, http.StatusInternalServerError, "failed to remove cart item", s.logger)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
