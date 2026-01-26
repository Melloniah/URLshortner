package model

import "time"

// Link represents a shortened URL in our system
type Link struct {
	ShortCode    string // the new link created
	OriginalURL  string 
	Clicks       int // how many times it has been accessed
	CreatedAt    time.Time  //time it was created
	LastAccessed time.Time //last time it was accessed
}

