package errors

import (
	"testing"
)

func TestDB_Error(t *testing.T) {
	err := &DB{}
	expected := "Database error"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestRowsEffectedError_Error(t *testing.T) {
	err := &RowsEffectedError{}
	expected := "No Rows Affected"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}
