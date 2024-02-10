package handlers

import "github.com/ASaidOguz/linkShortener/internal/application/interfaces"

// HTTPHandler provides HTTP handlers for the link shortener service
type HTTPHandler struct {
	Service interfaces.ShortenerService
}
