package auth

import (
	"testing"
	"net/http"
	"fmt"
)

func TestAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatalf("Coudldn't create a request")
	}

	key := "MY_API_KEY_HERE"
	req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", key))

	val, err := GetAPIKey(req.Header)

	if err != nil {
		t.Fatalf("Something went wrong: %s", err)
	}

	if val != key {
		t.Fatalf("expected: %v, got: %v", key, val)
	}
}
