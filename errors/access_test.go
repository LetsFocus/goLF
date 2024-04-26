package errors

import "testing"

func TestUnauthorized_Error(t *testing.T) {
	err := &Unauthorized{}
	expected := "Unauthorized: access denied"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestForbidden_Error(t *testing.T) {
	err := &Forbidden{}
	expected := "Forbidden: access forbidden"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}
