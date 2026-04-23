package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

func writeJSONError(w http.ResponseWriter, statusCode int, message string, logger *slog.Logger) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	response := entity.ErrorResponse{Error: message}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("failed to encode error response", err)
	}
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}, logger *slog.Logger) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("failed to encode response", err)
	}
}
