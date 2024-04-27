package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalError_Error(t *testing.T) {
	errMsg := "example error"
	err := &UnmarshalError{Err: errors.New(errMsg)}
	expected := fmt.Sprintf("Error while Unmarshalling Data, Error: '%s'", errMsg)
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestMarshalError_Error(t *testing.T) {
	errMsg := "example error"
	err := &MarshalError{Err: errors.New(errMsg)}
	expected := fmt.Sprintf("Error while Marshalling Data, Error: '%s'", errMsg)
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}
