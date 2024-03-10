package goLF

import (
	"mime/multipart"
	"net/http"
)

type Request struct {
	req *http.Request
}

type Requester interface {
	GetParam(string) string
	Bind(interface{}) error
	GetParams() map[string]interface{}
	GetPathParam() string
	GetHeader(string) string
	GetHeaders() map[string]interface{}
	MultiPartForm() (*multipart.Form, error)
	FormFile() (*multipart.FileHeader, error)
	BindYaml(any) error
	GetAuth() string
}
