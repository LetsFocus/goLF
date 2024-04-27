package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowed_Error(t *testing.T) {
	err := &MethodNotAllowed{}
	expected := "Method not allowed"
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}
