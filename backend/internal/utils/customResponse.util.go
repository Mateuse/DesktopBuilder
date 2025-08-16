package utils

import (
	"encoding/json"
	"net/http"

	"github.com/mateuse/desktop-builder-backend/internal/models"
)

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func WriteSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	resp := models.SuccessResponse{
		Code:    status,
		Message: message,
		Data:    data,
	}
	WriteJSON(w, status, resp)
}

func WriteError(w http.ResponseWriter, status int, message string, data interface{}) {
	resp := models.ErrorResponse{
		Code:    status,
		Message: message,
		Data:    data,
	}
	WriteJSON(w, status, resp)
}
