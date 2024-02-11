package services

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/ASaidOguz/linkShortener/internal/domain/entity"
)

// ShortenURL generates a short key for the original URL and saves it
func (s *ShortenerServiceImpl) ShortenURL(originalURL string) (*entity.ShortenedURL, error) {
	OriginalURL := &entity.URL{Original: originalURL}
	// set validation terms for simple url-checking...
	// solidity style XD
	success := OriginalURL.ValidateURL()
	if !success {
		return nil, errors.New("invalidated-url")
	}
	// create random shortkey for our link
	shortKey := s.generateShortKey()
	shortenedURL := &entity.ShortenedURL{
		OriginalURL: OriginalURL,
		ShortKey:    shortKey,
	}
	// save our shortenedURL into database
	err := s.Repository.Save(shortenedURL)
	if err != nil {
		return nil, err
	}
	return shortenedURL, nil
}

// Redirect retrieves the original URL by its short key and performs a redirect
func (s *ShortenerServiceImpl) Redirect(shortKey string) (string, error) {
	shortenedURL, err := s.Repository.FindByShortKey(shortKey)
	if err != nil {
		return "", err
	}
	return shortenedURL.OriginalURL.Original, nil
}

// generateShortKey generates a random short key for a URL
func (s *ShortenerServiceImpl) generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var shortKey strings.Builder
	for i := 0; i < 6; i++ {
		shortKey.WriteByte(charset[rand.Intn(len(charset))])
	}
	return shortKey.String()
}
