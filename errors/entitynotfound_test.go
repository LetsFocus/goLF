package errors

import (
	"fmt"
	"testing"
)

func TestEntityNotFound_Error(t *testing.T) {
	entity := "User"
	value := "123"
	err := &EntityNotFound{Entity: entity, Value: value}
	expected := fmt.Sprintf("No '%v' found for Id: '%v'", entity, value)
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}
