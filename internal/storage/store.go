package store

import "urlshortener/internal/model"

type Store interface {
	Create(link model.Link) error //save it somewhere, in this case model link
	Get(shortCode string) (model.Link, error) //retrieve it
	IncrementClicks(shortCode string) error //show the number of time its accesed
}