package model

import "time"

type Link struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortCode    string    `json:"short_code"`
	Clicks       int       `json:"clicks"`
	CreatedAt    time.Time `json:"created_at"`
	LastAccessed time.Time `json:"last_accessed,omitempty"`
}
