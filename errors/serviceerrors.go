package errors

import "fmt"

// ServiceCall represents a 500 Internal Server Error indicating an error in a service call.
type ServiceCall struct {
	Service string
	Err     error
}

func (e *ServiceCall) Error() string {
	return fmt.Sprintf("Error while calling %s service, Error: %v", e.Service, e.Err.Error())
}

// ServiceUnavailable represents a 500 Internal Server Error indicating a service is unavailable.
type ServiceUnavailable struct {
	Service string
}

func (e *ServiceUnavailable) Error() string {
	return fmt.Sprintf("Service '%s' is unavailable", e.Service)
}
