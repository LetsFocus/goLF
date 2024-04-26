package errors

// MethodNotAllowed represents a 405 Method Not Allowed error.
type MethodNotAllowed struct{}

func (e *MethodNotAllowed) Error() string {
	return "Method not allowed"
}
