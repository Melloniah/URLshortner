
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"urlshortener/internal/model"
	"urlshortener/internal/storage"
	"urlshortener/pkg/shortener"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// URLHandler handles URL shortening requests
type URLHandler struct {
	storage   *storage.MemoryStore
	generator *shortener.Generator
	baseURL   string
}

// NewURLHandler creates a new URL handler
func NewURLHandler(store *storage.MemoryStore, generator *shortener.Generator, baseURL string) *URLHandler {
	return &URLHandler{
		storage:   store,
		generator: generator,
		baseURL:   baseURL,
	}
}

// CreateShortURLRequest is the request body
type CreateShortURLRequest struct {
	URL string `json:"url"`
}

// CreateShortURLResponse is the response
type CreateShortURLResponse struct {
	ShortCode string `json:"short_code"`
	ShortURL  string `json:"short_url"`
	LongURL   string `json:"long_url"`
}

// CreateShortURL handles POST /api/shorten
func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req CreateShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := h.generator.Generate()

	link := model.Link{
		ID:          uuid.New().String(),
		OriginalURL: req.URL,
		ShortCode:   shortCode,
		Clicks:      0,
		CreatedAt:   time.Now(),
	}

	if err := h.storage.Create(link); err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	response := CreateShortURLResponse{
		ShortCode: shortCode,
		ShortURL:  h.baseURL + "/" + shortCode,
		LongURL:   req.URL,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// RedirectURL handles GET /{shortCode}
func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	link, err := h.storage.Get(shortCode)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// This increments the clicks in memory and saves to links.json
	h.storage.IncrementClicks(shortCode)

	http.Redirect(w, r, link.OriginalURL, http.StatusFound)
}

// GetStats handles GET /api/stats/{shortCode}
func (h *URLHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	link, err := h.storage.Get(shortCode)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}
