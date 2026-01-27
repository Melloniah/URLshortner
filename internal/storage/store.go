

package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"urlshortener/internal/model"
)

const dataFile = "links.json"

// MemoryStore stores links in memory and persists to file
type MemoryStore struct {
	mu    sync.Mutex
	links map[string]model.Link
}

// NewMemoryStore creates a new storage and loads existing data
func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{
		links: make(map[string]model.Link),
	}
	
	// Load existing data from file
	store.loadFromFile()
	
	return store
}

// loadFromFile reads links from JSON file
func (s *MemoryStore) loadFromFile() error {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		// File doesn't exist yet, that's okay
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var links []model.Link
	if err := json.Unmarshal(file, &links); err != nil {
		return err
	}

	// Load into map
	for _, link := range links {
		s.links[link.ShortCode] = link
	}

	return nil
}

// saveToFile writes links to JSON file
func (s *MemoryStore) saveToFile() error {
	// Convert map to slice
	links := make([]model.Link, 0, len(s.links))
	for _, link := range s.links {
		links = append(links, link)
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(dataFile, data, 0644)
}

// Create saves a new link
func (s *MemoryStore) Create(link model.Link) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.links[link.ShortCode]; exists {
		return errors.New("short code already exists")
	}

	s.links[link.ShortCode] = link
	
	// Persist to file
	return s.saveToFile()
}

// Get finds a link by its short code
func (s *MemoryStore) Get(shortCode string) (model.Link, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	link, exists := s.links[shortCode]
	if !exists {
		return model.Link{}, errors.New("link not found")
	}

	return link, nil
}

// IncrementClicks increases click count and updates LastAccessed
func (s *MemoryStore) IncrementClicks(shortCode string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	link, exists := s.links[shortCode]
	if !exists {
		return errors.New("link not found")
	}

	link.Clicks++
	s.links[shortCode] = link
	
	// Persist to file
	return s.saveToFile()
}

// GetAll returns all links (useful for stats/admin)
func (s *MemoryStore) GetAll() []model.Link {
	s.mu.Lock()
	defer s.mu.Unlock()

	links := make([]model.Link, 0, len(s.links))
	for _, link := range s.links {
		links = append(links, link)
	}

	return links
}