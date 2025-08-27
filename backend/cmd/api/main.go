package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mateuse/desktop-builder-backend/internal/routes"
	"github.com/mateuse/desktop-builder-backend/internal/utils"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get allowed origins from environment variable or use defaults
		allowedOriginsStr := os.Getenv("CORS_ALLOWED_ORIGINS")
		if allowedOriginsStr == "" {
			allowedOriginsStr = "http://localhost:3000,http://localhost:3001,http://localhost:3002,http://localhost:5173"
		}

		// Parse allowed origins into a slice
		allowedOrigins := strings.Split(allowedOriginsStr, ",")

		// Get the origin from the request
		requestOrigin := r.Header.Get("Origin")

		// Check if the request origin is in our allowed list
		originAllowed := false
		for _, origin := range allowedOrigins {
			if strings.TrimSpace(origin) == requestOrigin {
				originAllowed = true
				break
			}
		}

		// Set CORS headers
		if originAllowed {
			w.Header().Set("Access-Control-Allow-Origin", requestOrigin)
		}
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize database connection
	if err := utils.InitializeDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := utils.CloseDatabase(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Initialize Redis connection
	if err := utils.InitializeRedis(); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	defer func() {
		if err := utils.CloseRedis(); err != nil {
			log.Printf("Error closing Redis: %v", err)
		}
	}()

	mux := http.NewServeMux()
	routes.RegisterHealthRoutes(mux)
	routes.RegisterComponentRoutes(mux)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Backend is running on port %s", port)
	if err := http.ListenAndServe(":"+port, withCORS(mux)); err != nil {
		log.Fatal(err)
	}
}
