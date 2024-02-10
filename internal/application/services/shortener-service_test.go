package services

import (
	"errors"
	"testing"

	"github.com/ASaidOguz/linkShortener/internal/domain/entity"
	"github.com/golang/mock/gomock"
)

// MockRepository is a mock implementation of ShortenedURLRepository for testing
type MockRepository struct {
	mockCtrl *gomock.Controller
}

// Save mocks the Save method of ShortenedURLRepository
func (m *MockRepository) Save(shortenedURL *entity.ShortenedURL) error {
	// Mock implementation
	return nil
}

// FindByShortKey mocks the FindByShortKey method of ShortenedURLRepository
func (m *MockRepository) FindByShortKey(shortKey string) (*entity.ShortenedURL, error) {
	// Mock implementation
	if shortKey == "existingKey" {
		return &entity.ShortenedURL{
			OriginalURL: &entity.URL{Original: "https://example.com"},
			ShortKey:    "existingKey",
		}, nil
	}
	return nil, errors.New("Short URL not found")
}

func TestShortenURL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := MockRepository{mockCtrl}
	service := ShortenerServiceImpl{Repository: &mockRepo}

	// Test case: Successfully shorten URL
	shortenedURL, err := service.ShortenURL("https://example.com")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if shortenedURL == nil {
		t.Error("Expected shortenedURL to be non-nil, got nil")
	}

}

func TestRedirect(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := MockRepository{mockCtrl}
	service := ShortenerServiceImpl{Repository: &mockRepo}

	// Test case: Redirect to existing short key
	redirectURL, err := service.Redirect("existingKey")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedURL := "https://example.com"
	if redirectURL != expectedURL {
		t.Errorf("Expected redirect URL: %s, got: %s", expectedURL, redirectURL)
	}

	// Test case: Redirect to non-existing short key
	redirectURL, err = service.Redirect("nonExistingKey")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if redirectURL != "" {
		t.Errorf("Expected redirect URL to be empty, got: %s", redirectURL)
	}

	// Additional test cases can be added to cover edge cases and error scenarios
}
