package store

import (
	"errors"
	"sync"
     "time"
	"urlshortener/internal/model"

)
// mu locks the store so two requests don’t modify data at the same time
type MemoryStore struct{
	mu 	sync.Mutex
	links map[string]model.Link
}

// creates a new empty storage
func NewMemoryStore () *MemoryStore{ 
	return &MemoryStore{
		links:make(map[string]model.Link)	}
}

// Saves a new link, returns error if short code already exists
func (s *MemoryStore) Create(link model.Link) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.links[link.ShortCode]; exists {
		return errors.New("short code already exists")
	}

	s.links[link.ShortCode] = link
	return nil
}

//implement get, which returns zero value and error if its failling
// mutexes help with concurrent access. Finds a link by its short code
 func (s *MemoryStore) Get(shortCode string) (model.Link, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	link, exists := s.links[shortCode]
	if !exists {
		return model.Link{}, errors.New("link not found")
	}

	return link, nil
}
 //implements increment clicks
 func (s *MemoryStore) IncrementClicks(shortCode string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	link, exists := s.links[shortCode]
	if !exists {
		return errors.New("link not found")
	}

	link.Clicks++
	link.LastAccessed = time.Now()
	s.links[shortCode] = link

	return nil
}
