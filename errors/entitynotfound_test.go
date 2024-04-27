package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityNotFound_Error(t *testing.T) {
	entity := "User"
	value := "123"
	err := &EntityNotFound{Entity: entity, Value: value}
	expected := fmt.Sprintf("No '%v' found for Id: '%v'", entity, value)
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}
