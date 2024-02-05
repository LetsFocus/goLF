package configs

import (
	"github.com/LetsFocus/goLF/logger"
	"strings"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name       string
		env        map[string]string
		expected   string
		expectLogs string
	}{
		{
			name:     "Load .env file successfully",
			env:      map[string]string{"APP_ENV": "test"},
			expected: "./configs/.test.env",
		},
		{
			name:     "No APP_ENV set, load default .env",
			env:      map[string]string{},
			expected: "./configs/.env",
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			log := logger.NewCustomLogger()

			for key, value := range test.env {
				t.Setenv(key, value)
			}

			config := NewConfig(log)

			if !strings.Contains(config.Log.GetLog(), test.expected) {
				t.Errorf("Test Case Failed[%v] Got: %v, Required: %v", i+1, config.Log.GetLog(), test.expected)
			}
		})
	}
}

func TestConfig_Get(t *testing.T) {
	tests := []struct {
		name     string
		env      map[string]string
		key      string
		expected string
	}{
		{
			name:     "Get value from environment variable",
			env:      map[string]string{"MY_VARIABLE": "test_value"},
			key:      "MY_VARIABLE",
			expected: "test_value",
		},
		{
			name:     "Variable not found, return empty string",
			env:      map[string]string{},
			key:      "NOT_FOUND_VARIABLE",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.env {
				t.Setenv(key, value)
			}

			config := Config{}
			result := config.Get(test.key)

			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
