package infrastructure

import (
	"fmt"

	"github.com/ASaidOguz/linkShortener/internal/domain/entity"
)

// InMemoryRepository is an in-memory implementation of ShortenedURLRepository
type InMemoryRepository struct {
	UrlMap map[string]*entity.ShortenedURL
}

// Save saves a ShortenedURL object to the in-memory repository
func (r *InMemoryRepository) Save(shortenedURL *entity.ShortenedURL) error {
	r.UrlMap[shortenedURL.ShortKey] = shortenedURL
	return nil
}

// FindByShortKey retrieves a ShortenedURL object by its short key from the in-memory repository
func (r *InMemoryRepository) FindByShortKey(shortKey string) (*entity.ShortenedURL, error) {
	shortenedURL, ok := r.UrlMap[shortKey]
	if !ok {
		return nil, fmt.Errorf("short url not found")
	}
	return shortenedURL, nil
}
