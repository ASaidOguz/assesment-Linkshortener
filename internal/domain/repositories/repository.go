package repositories

import (
	"github.com/ASaidOguz/linkShortener/internal/domain/entity"
)

// ShortenedURLRepository is an interface for storing and retrieving ShortenedURL objects
type ShortenedURLRepository interface {
	Save(shortenedURL *entity.ShortenedURL) error
	FindByShortKey(shortKey string) (*entity.ShortenedURL, error)
}
