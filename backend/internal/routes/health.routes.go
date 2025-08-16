package routes

import (
	"net/http"

	"github.com/mateuse/desktop-builder-backend/internal/handlers"
)

func RegisterHealthRoutes(router *http.ServeMux) {
	router.HandleFunc("/health", handlers.HealthHandler)
}
