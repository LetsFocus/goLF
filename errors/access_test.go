package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnauthorized_Error(t *testing.T) {
	err := &Unauthorized{}
	expected := "Unauthorized: access denied"
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestForbidden_Error(t *testing.T) {
	err := &Forbidden{}
	expected := "Forbidden: access forbidden"
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}
