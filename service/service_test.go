package service

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNewClient(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		input        *logger.CustomLogger
// 		resourceAddr string
// 		expectedURL  string
// 		log          string
// 	}{
// 		{
// 			name:         "Empty resource address",
// 			input:        logger.NewCustomLogger(),
// 			resourceAddr: "",
// 			expectedURL:  "",
// 			log:          "value for resourceAddress is empty",
// 		},
// 		{
// 			name:         "Non-empty resource address",
// 			input:        logger.NewCustomLogger(),
// 			resourceAddr: "http://example.com",
// 			expectedURL:  "http://example.com",
// 			log:          "",
// 		},
// 	}

// 	for i, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			client := NewClient(tc.resourceAddr, tc.input)

// 			if strings.Contains(tc.input.GetLog(), tc.log) {
// 				t.Errorf("Testcase Failed[%v], Required Log: %v, Got: %v", i+1, tc.input.GetLog(), tc.log)
// 			}

// 			assert.NotNil(t, client)
// 			assert.Equal(t, tc.expectedURL, client.url)
// 		})
// 	}
// }

func TestClient_createRequest(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		target         string
		params         map[string]interface{}
		body           []byte
		headers        map[string]string
		expectedMethod string
		expectedURI    string
	}{
		{
			name:           "Empty Target",
			method:         http.MethodGet,
			target:         "",
			params:         nil,
			body:           nil,
			headers:        nil,
			expectedMethod: http.MethodGet,
			expectedURI:    "http://example.com",
		},
		{
			name:           "Non-empty Target",
			method:         http.MethodPost,
			target:         "path/to/resource",
			params:         map[string]interface{}{"param1": "value1"},
			body:           []byte(`{"key":"value"}`),
			headers:        map[string]string{"Authorization": "Bearer token"},
			expectedMethod: http.MethodPost,
			expectedURI:    "http://example.com/path/to/resource?param1=value1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := Client{url: "http://example.com"}
			req, err := client.createRequest(context.Background(), tc.method, tc.target, tc.params, tc.body, tc.headers)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedMethod, req.Method)
			assert.Equal(t, tc.expectedURI, req.URL.String())
		})
	}
}

