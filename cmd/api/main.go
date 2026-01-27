
package main

import (
	"fmt"
	"log"
	"net/http"

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
	baseURL := "http://localhost:8080"

	// Initialize handler
	handler := handlers.NewURLHandler(store, generator, baseURL)

	// Setup router
	router := mux.NewRouter()

	// API routes
	router.HandleFunc("/api/shorten", handler.CreateShortURL).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stats/{shortCode}", handler.GetStats).Methods("GET", "OPTIONS")

	// Redirect route
	router.HandleFunc("/{shortCode}", handler.RedirectURL).Methods("GET")

	// Setup CORS with more permissive settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Start server
	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("React frontend: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}