package configs

import (
	"testing"
)

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
