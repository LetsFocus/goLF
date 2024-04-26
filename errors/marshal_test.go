package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnmarshalError_Error(t *testing.T) {
	errMsg := "example error"
	err := &UnmarshalError{Err: errors.New(errMsg)}
	expected := fmt.Sprintf("Error while Unmarshalling Data, Error: '%s'", errMsg)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestMarshalError_Error(t *testing.T) {
	errMsg := "example error"
	err := &MarshalError{Err: errors.New(errMsg)}
	expected := fmt.Sprintf("Error while Marshalling Data, Error: '%s'", errMsg)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}
