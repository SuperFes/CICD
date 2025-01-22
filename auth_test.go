package main_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Log("TestGetAPIKey")

	headers := http.Header{}

	headers.Add("Authorization", "Bearer 123456")

	_, err := auth.GetAPIKey(headers)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err.Error() != "malformed authorization header" {
		t.Errorf("Expected error, got %s", err.Error())
	}
}
