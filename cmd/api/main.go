
package main

import (
	"fmt"
	"log"
	"net/http"

	"urlshortener/internal/handlers"
	"urlshortener/internal/storage"
	"urlshortener/pkg/shortener"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/api/shorten", handler.CreateShortURL).Methods("POST")
	router.HandleFunc("/api/stats/{shortCode}", handler.GetStats).Methods("GET")

	// Redirect route
	router.HandleFunc("/{shortCode}", handler.RedirectURL).Methods("GET")

	// Start server
	fmt.Println("🚀 Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
