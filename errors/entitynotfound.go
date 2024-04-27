package errors

import "fmt"

// EntityNotFound represents a 404 Not Found error indicating an entity was not found.
type EntityNotFound struct {
	Entity string
	Value  string
}

func (e *EntityNotFound) Error() string {
	return fmt.Sprintf("No '%v' found for Id: '%v'", e.Entity, e.Value)
}
