package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestServiceCall_Error(t *testing.T) {
	service := "exampleService"
	errMsg := "example error"
	err := &ServiceCall{Service: service, Err: errors.New(errMsg)}
	expected := fmt.Sprintf("Error while calling %s service, Error: %s", service, errMsg)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestServiceUnavailable_Error(t *testing.T) {
	service := "exampleService"
	err := &ServiceUnavailable{Service: service}
	expected := fmt.Sprintf("Service '%s' is unavailable", service)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}
