
package storage

import (
	"errors"
	"sync"
	"time"
	"urlshortener/internal/model"
)

// MemoryStore stores links in memory
type MemoryStore struct {
	mu    sync.Mutex
	links map[string]model.Link
}

// NewMemoryStore creates a new empty storage
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		links: make(map[string]model.Link),
	}
}

// Create saves a new link
func (s *MemoryStore) Create(link model.Link) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.links[link.ShortCode]; exists {
		return errors.New("short code already exists")
	}

	s.links[link.ShortCode] = link
	return nil
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
	link.LastAccessed = time.Now()  // Track when it was accessed
	s.links[shortCode] = link
	return nil
}
