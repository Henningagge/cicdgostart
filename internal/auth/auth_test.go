package auth

import (
	"net/http"
	"testing"
)

func testApi(t *testing.T){
	testCases := []struct {
		name string
		headers http.Header
		expectedAPIKey  string
		exPectedError error
	}{
		{
			name: "Valid Key",
			headers: http.Header{"Authorization": {"ApiKey 12345"}},
			expectedAPIKey: "dskfsalkfdlsak", 
			exPectedError: nil,
		},
	}
	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T){
			ApiKey, err := GetAPIKey(tc.headers)
			if err != tc.exPectedError{
				t.Errorf("Expected error: %v, got: %v", tc.exPectedError, err)
			}
			if ApiKey != tc.expectedAPIKey{
				t.Errorf("Expected API Key: %s, got: %s", tc.expectedAPIKey, ApiKey)
			}
		})
	}
}