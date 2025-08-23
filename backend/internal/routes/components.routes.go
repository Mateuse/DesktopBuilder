package routes

import (
	"net/http"

	"github.com/mateuse/desktop-builder-backend/internal/handlers"
)

func RegisterComponentRoutes(router *http.ServeMux) {
	router.HandleFunc("/components", handlers.GetComponentsHandler)
	router.HandleFunc("/components/{category}", handlers.GetComponentsHandler)
	router.HandleFunc("/components/{category}/{brand}", handlers.GetComponentsHandler)
	router.HandleFunc("/components/item/{id}", handlers.GetComponentsHandler)
}
