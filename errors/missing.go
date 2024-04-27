package errors

import (
	"fmt"
	"strings"
)

// MissingHeader represents a 400 Bad Request error indicating a missing header.
type MissingHeader struct {
	Header string
}

func (e *MissingHeader) Error() string {
	return fmt.Sprintf("Missing header: '%s'", e.Header)
}

// MissingHeaders represents a 400 Bad Request error indicating multiple missing headers.
type MissingHeaders struct {
	Headers []string
}

func (e *MissingHeaders) Error() string {
	return fmt.Sprintf("Missing headers: '%s'", strings.Join(e.Headers, "', '"))
}

// MissingParam represents a 400 Bad Request error indicating a missing parameter.
type MissingParam struct {
	Param string
}

func (e *MissingParam) Error() string {
	return fmt.Sprintf("Missing parameter: '%s'", e.Param)
}

// MissingParams represents a 400 Bad Request error indicating multiple missing parameters.
type MissingParams struct {
	Params []string
}

func (e *MissingParams) Error() string {
	return fmt.Sprintf("Missing parameters: '%s'", strings.Join(e.Params, "', '"))
}
