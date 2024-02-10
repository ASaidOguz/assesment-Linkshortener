package handlers

import (
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
	return &entity.ShortenedURL{
		OriginalURL: &entity.URL{Original: "https://example.com"},
		ShortKey:    "testkey",
	}, nil

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

	// Add more assertions as needed
}

// Add more test cases to cover other scenarios, such as invalid input and error conditions
