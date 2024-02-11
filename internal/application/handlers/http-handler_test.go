package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ASaidOguz/linkShortener/internal/domain/entity"
)

type MockService struct {
}

// This is mock ShortenURL to satify the interface and get mock constant values...
func (s *MockService) ShortenURL(originalURL string) (*entity.ShortenedURL, error) {
	if originalURL == "https://example.com" {
		return &entity.ShortenedURL{
			OriginalURL: &entity.URL{Original: "https://example.com"},
			ShortKey:    "testkey",
		}, nil
	}
	return nil, errors.New("invalidated-url")

}

// This is mock Redirect to satify the interface and get mock constant values...
func (s *MockService) Redirect(shortKey string) (string, error) {

	return "https://example.com", nil
}
func TestShortenURLHandler_ValidInput(t *testing.T) {

	service := &MockService{}

	handler := &HTTPHandler{
		Service: service,
	}

	// Prepare a request with valid input
	reqBody := "url=https://example.com"
	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	handler.ShortenURLHandler(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expectedShortURL := "http://localhost:8080/testkey"
	if !strings.Contains(w.Body.String(), expectedShortURL) {
		t.Errorf("Expected response body to contain %q, got %q", expectedShortURL, w.Body.String())
	}

}

func TestShortenURLHandler_InvalidInput(t *testing.T) {
	service := &MockService{}

	handler := &HTTPHandler{
		Service: service,
	}

	// Prepare a request with valid input
	reqBody := "url=example.com"
	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	handler.ShortenURLHandler(w, req)

	// Check the response status code
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestRedirectURLHandler_Works(t *testing.T) {
	// Prepare a mock service
	service := &MockService{}

	// Create a handler
	handler := &HTTPHandler{
		Service: service,
	}

	// Prepare a request with a valid short key
	req, err := http.NewRequest("GET", "/testkey", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Call the RedirectHandler function
	handler.RedirectHandler(w, req)

	// Check the response status code
	fmt.Println("response status code:", w.Code)
	if w.Code != http.StatusFound {
		t.Errorf("Expected status code %d, got %d", http.StatusFound, w.Code)
	}

	// Check the redirected URL
	fmt.Println("Location get:", w.Header().Get("Location"))
	expectedURL := "https://example.com"
	if location := w.Header().Get("Location"); location != expectedURL {
		t.Errorf("Expected redirect to %s, got %s", expectedURL, location)
	}
}
