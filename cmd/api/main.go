
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"urlshortener/internal/handlers"
	"urlshortener/internal/storage"
	"urlshortener/pkg/shortener"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize storage and generator
	store := storage.NewMemoryStore()
	generator := shortener.NewGenerator()
	
	// Get base URL from environment or use default
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// Get allowed origins from environment
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:3000"
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize handler
	handler := handlers.NewURLHandler(store, generator, baseURL)

	// Setup router
	router := mux.NewRouter()

	// API routes
	router.HandleFunc("/api/shorten", handler.CreateShortURL).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stats/{shortCode}", handler.GetStats).Methods("GET", "OPTIONS")

	// Redirect route
	router.HandleFunc("/{shortCode}", handler.RedirectURL).Methods("GET")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{allowedOrigin, "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Start server
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Printf("Allowed origin: %s\n", allowedOrigin)
	fmt.Printf("Base URL: %s\n", baseURL)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(router)))
}
