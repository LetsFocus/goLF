// package service

// import (
// 	"testing"

// 	"github.com/LetsFocus/goLF/logger"
// 	"github.com/stretchr/testify/assert"
// )

// func TestNewClient(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		resourceAddr string
// 		expectedURL  string
// 		expectLog    bool
// 	}{
// 		{"Empty resource address", "", "", true},
// 		{"Non-empty resource address", "http://example.com", "http://example.com", false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockLogger := &logger.CustomLogger{}
// 			client := NewClient(tt.resourceAddr, mockLogger)

// 			if tt.expectLog {
// 				assert.Contains(t, mockLogger.Logs(), "value for resourceAddress is empty")
// 			} else {
// 				assert.NotContains(t, mockLogger.Logs(), "value for resourceAddress is empty")
// 			}

// 			assert.NotNil(t, client)
// 			assert.Equal(t, tt.expectedURL, client.url)
// 		})
// 	}
// }
