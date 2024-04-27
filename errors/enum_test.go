package errors

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumError_Error(t *testing.T) {
	supportedValues := []string{"A", "B", "C"}
	value := "D"
	err := &EnumError{SupportedValues: supportedValues, Value: value}
	expected := fmt.Sprintf("Value of the field %s must be one of [%s]", value, strings.Join(supportedValues, ", "))
	assert.Equal(t, expected, err.Error(), "Expected error message '%s', got '%s'", expected, err.Error())
}
