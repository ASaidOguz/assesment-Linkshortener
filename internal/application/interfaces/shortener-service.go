package interfaces

import (
	"github.com/ASaidOguz/linkShortener/internal/domain/entity"
)

// ShortenerService provides methods for shortening and redirecting URLs
type ShortenerService interface {
	ShortenURL(originalURL string) (*entity.ShortenedURL, error)
	Redirect(shortKey string) (string, error)
}
