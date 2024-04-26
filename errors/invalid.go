package errors


// InvalidParam represents a 400 Bad Request error indicating an invalid parameter.
type InvalidParam struct {
	Param string
}

func (e *InvalidParam) Error() string {
	return fmt.Sprintf("Invalid parameter: '%s'", e.Param)
}

// InvalidParams represents a 400 Bad Request error indicating multiple invalid parameters.
type InvalidParams struct {
	Params []string
}

func (e *InvalidParams) Error() string {
	return fmt.Sprintf("Invalid parameters: '%s'", strings.Join(e.Params, "', '"))
}

// InvalidHeader represents a 400 Bad Request error indicating an invalid header.
type InvalidHeader struct {
	Header string
}

func (e *InvalidHeader) Error() string {
	return fmt.Sprintf("Invalid header: '%s'", e.Header)
}

// InvalidHeaders represents a 400 Bad Request error indicating multiple invalid headers.
type InvalidHeaders struct {
	Headers []string
}

func (e *InvalidHeaders) Error() string {
	return fmt.Sprintf("Invalid headers: '%s'", strings.Join(e.Headers, "', '"))
}
