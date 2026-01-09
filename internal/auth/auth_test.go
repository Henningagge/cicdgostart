package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name           string
		headers        http.Header
		expectedAPIKey string
		expectedError  error
	}{
		{
			name:           "Valid API Ke",
			headers:        http.Header{"Authorization": {"ApiKey 12345"}},
			expectedAPIKey: "12345",
			expectedError:  nil,
		},
		{
			name:           "No Authorization Header",
			headers:        http.Header{},
			expectedAPIKey: "",
			expectedError:  ErrNoAuthHeaderIncluded,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)

			if err != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			if apiKey != tc.expectedAPIKey {
				t.Errorf("Expected API Key: %s, got: %s", tc.expectedAPIKey, apiKey)
			}
		})
	}
}
