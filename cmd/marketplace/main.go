package main

import (
	"fmt"
	"net/http"

	"github.com/ASaidOguz/linkShortener/internal/application/handlers"
	"github.com/ASaidOguz/linkShortener/internal/application/services"
	"github.com/ASaidOguz/linkShortener/internal/domain/entity"
	"github.com/ASaidOguz/linkShortener/internal/infrastructure"
)

func main() {
	//currenlty used port
	port := ":8080"

	// Initialize repository
	repo := &infrastructure.InMemoryRepository{
		UrlMap: make(map[string]*entity.ShortenedURL),
	}

	// Initialize service
	service := &services.ShortenerServiceImpl{
		Repository: repo,
	}

	// Initialize HTTP handler
	handler := &handlers.HTTPHandler{
		Service: service,
	}

	// Define HTTP routes
	http.HandleFunc("/shorten", handler.ShortenURLHandler)
	http.HandleFunc("/", handler.RedirectHandler)

	// Start the HTTP server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(port, nil)
}
