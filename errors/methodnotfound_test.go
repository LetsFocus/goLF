package errors

import (
	"testing"
)

func TestMethodNotAllowed_Error(t *testing.T) {
	err := &MethodNotAllowed{}
	expected := "Method not allowed"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}
