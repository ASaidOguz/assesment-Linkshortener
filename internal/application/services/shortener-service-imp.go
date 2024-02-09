package services

import "github.com/ASaidOguz/linkShortener/internal/domain/repositories"

// ShortenerServiceImpl is an implementation of ShortenerService
type ShortenerServiceImpl struct {
	Repository repositories.ShortenedURLRepository
}
