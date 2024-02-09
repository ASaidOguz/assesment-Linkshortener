package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ASaidOguz/linkShortener/internal/application/services"
)

// HTTPHandler provides HTTP handlers for the link shortener service
type HTTPHandler struct {
	Service services.ShortenerServiceImpl
}

// ShortenURLHandler handles the request to shorten a URL
func (h *HTTPHandler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortenedURL, err := h.Service.ShortenURL(originalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortenedURL.ShortKey)
	fmt.Fprintf(w, "Shortened URL: %s\n", shortURL)
}

// RedirectHandler handles the request to redirect to the original URL
func (h *HTTPHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/")
	originalURL, err := h.Service.Redirect(shortKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}
