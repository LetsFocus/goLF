package errors

import (
	"fmt"
	"strings"
)

// EnumError represents a 400 Bad Request error indicating an invalid enumeration value.
type EnumError struct {
	SupportedValues []string
	Value           string
}

func (e *EnumError) Error() string {
	return fmt.Sprintf("Value of the field %s must be one of [%s]", e.Value, strings.Join(e.SupportedValues, ", "))
}
