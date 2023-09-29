package goLF

type Response struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Code       string      `json:"code,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
