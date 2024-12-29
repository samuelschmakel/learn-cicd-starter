package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"ApiKey Expected"},
		"Accept":        {"application/json", "text/html"},
	}

	expected := "Expected"

	result, err := GetAPIKey(header)
	if err != nil {
		t.Errorf("Got an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetAPIKeyNoApiKey(t *testing.T) {
	header := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer Token"},
		"Accept":        {"application/json", "text/html"},
	}

	expectedErr := "malformed authorization header"

	_, err := GetAPIKey(header)
	if err == nil || err.Error() != expectedErr {
		t.Errorf("Expected %v, got %v", errors.New("malformed authorization header"), err)
	}
}
