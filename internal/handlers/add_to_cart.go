package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cpbartem2158/CART_API/internal/entity"
	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (s *Server) AddCartItemToCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed", s.logger)
		return
	}

	idStr := r.PathValue("id")
	cartId, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid cart ID", s.logger)
		return
	}

	var request entity.AddItemRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid request body", s.logger)
		return
	}
	if strings.TrimSpace(request.Product) == "" {
		writeJSONError(w, http.StatusBadRequest, "product name is required", s.logger)
		return
	}

	if request.Price <= 0 {
		writeJSONError(w, http.StatusBadRequest, "price must be positive", s.logger)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	cartItem, err := s.service.AddCartItemToCart(ctx, cartId, request.Product, request.Price)
	if err != nil {
		if errors.Is(err, errorsx.ErrCartNotFound) {
			writeJSONError(w, http.StatusNotFound, "cart not found", s.logger)
			return
		}
		if errors.Is(err, errorsx.ErrCartItemNotFound) {
			writeJSONError(w, http.StatusNotFound, "cart item not found", s.logger)
			return
		}
		if errors.Is(err, errorsx.ErrCartFull) {
			writeJSONError(w, http.StatusInternalServerError, "cart is full (maximum 5 items)", s.logger)
			return
		}
		s.logger.Error("error adding cartItem to cart", "error", err)
		writeJSONError(w, http.StatusInternalServerError, "error adding cartItem to cart", s.logger)
		return
	}
	writeJSONResponse(w, http.StatusOK, cartItem, s.logger)
}
