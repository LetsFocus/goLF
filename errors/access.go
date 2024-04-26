package errors

// Unauthorized represents a 401 Unauthorized error.
type Unauthorized struct{}

func (e *Unauthorized) Error() string {
	return "Unauthorized: access denied"
}

// Forbidden represents a 403 Forbidden error.
type Forbidden struct{}

func (e *Forbidden) Error() string {
	return "Forbidden: access forbidden"
}
