package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		expectedKey string
		expectError bool
	}{
		{
			name:        "valid api key",
			headerValue: "ApiKey abc123",
			expectedKey: "abc123",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			headerValue: "",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "wrong auth scheme",
			headerValue: "Bearer abc123",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "missing api key",
			headerValue: "ApiKey",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "empty api key",
			headerValue: "ApiKey",
			expectedKey: "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}

			if tc.headerValue != "" {
				headers.Set("Authorization", tc.headerValue)
			}

			key, err := GetAPIKey(headers)

			if tc.expectError && err == nil {
				t.Fatalf("expected an error, got nil")
			}

			if !tc.expectError && err != nil {
				t.Fatalf("did not expect an error, got %v", err)
			}

			if key != tc.expectedKey {
				t.Errorf("expected key %q, got %q", tc.expectedKey, key)
			}
		})
	}
}
