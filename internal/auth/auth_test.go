package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{}

	key, err := GetAPIKey(headers)

	if key != "" {
		t.Fatalf("expected empty key, got %q", key)
	}
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_Valid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key != "abc123" {
		t.Fatalf("expected key %q, got %q", "abc123", key)
	}
}
