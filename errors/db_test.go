package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_Error(t *testing.T) {
	err := &DB{}
	expected := "Database error"
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}

func TestRowsEffectedError_Error(t *testing.T) {
	err := &RowsEffectedError{}
	expected := "No Rows Affected"
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}
