package errors

import "fmt"

// UnmarshalError represents a 400 Bad Request error indicating a failure to unmarshal data.
type UnmarshalError struct {
	Err error
}

func (e *UnmarshalError) Error() string {
	return fmt.Sprintf("Error while Unmarshalling Data, Error: '%s'", e.Err.Error())
}

// MarshalError represents a 400 Bad Request error indicating a failure to marshal data.
type MarshalError struct {
	Err error
}

func (e *MarshalError) Error() string {
	return fmt.Sprintf("Error while Marshalling Data, Error: '%s'", e.Err.Error())
}
