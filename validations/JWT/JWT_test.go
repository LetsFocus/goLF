package JWT

import (
	"crypto/rsa"
	"net/http"
	"testing"

	"github.com/LetsFocus/goLF/errors"
)

func TestGetRSAPublicKey(t *testing.T) {
	tests := []struct {
		name           string
		publicKeyInput string
		expectedKey    *rsa.PublicKey
		expectedError  error
	}{
		{
			name:           "ValidPublicKey",
			publicKeyInput: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAy3W9nKct0qYDL4+xx8yK\n4k2GJk5iMKb6RgwQIb9h5oJvJv14Xp2gXm4PDLGnHxyS7N28pP2/7gdJ0CC+WbyZ\n9BoI57lkw6vVllPqGehL6hsxPPrNzzKx7H9qV2WqZqZbZbHcOKH96Vtiw2zCxItB\nC7++t5lZtjt+N/gGRnM8DZ0mG0m6p4VJcA7ZUI8Y2WSauIzjGtEh0QLUnATmh6qI\nN5AUKR0JcUYfF9vcDoRYlTxj8eT2+ZX7nR2Xz5UswrDB7o+XzRcIINQ8H8De7jDL\nGK0pBswMlTQXUBzF6adg+ztAaEMFCDVVl11Cg8utHtS99VjF9rJ6ff5rm7VVzYF8\nnwIDAQAB",
			expectedKey:    &rsa.PublicKey{},
			expectedError:  nil,
		},
		{
			name:           "InvalidPublicKey",
			publicKeyInput: "invalid public key",
			expectedKey:    nil,
			expectedError: &errors.Errors{
				StatusCode: http.StatusBadRequest,
				Code:       http.StatusText(http.StatusBadRequest),
				Reason:     "invalid PublicKey x509: malformed certificate",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetRSAPublicKey(tt.publicKeyInput)
			if err != nil {
				if tt.expectedError == nil || err.Error() != tt.expectedError.Error() {
					t.Errorf("Unexpected error: got %v, want %v", err, tt.expectedError)
				}
				return
			}
			if key == nil {
				t.Error("Expected key to be non-nil")
			}
		})
	}
}