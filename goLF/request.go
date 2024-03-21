package goLF

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	req *http.Request
}

type Requester interface {
	GetParam(string) string
	Bind(interface{}) error
	GetParams() map[string]string
	GetParamsArray() map[string][]string
	GetPathParam() string
	GetHeader(string) string
	GetHeaders() map[string][]string
	//MultiPartForm() (*multipart.Form, error)
	//FormFile() (*multipart.FileHeader, error)
	//BindYaml(any) error
	//GetAuth() string
}

func (r *Request) Bind(dataType interface{}) error {
	bodyBytes, err := io.ReadAll(r.req.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, &dataType)
}

func (r *Request) GetPathParam() string {
	return r.req.URL.Path
}

func (r *Request) GetHeader(key string) string {
	return r.req.Header.Get(key)
}

func (r *Request) GetHeaders() map[string][]string {
	return r.req.Header
}

func (r *Request) GetParam(key string) string {
	return r.req.URL.Query().Get(key)
}

func (r *Request) GetParams() map[string]string {
	query := r.req.URL.Query()
	queryParams := make(map[string]string)

	for key, value := range query {
		queryParams[key] = strings.Join(value, ",")
	}

	return queryParams
}

func (r *Request) GetParamsArray() map[string][]string {
	query := r.req.URL.Query()
	var queryParams map[string][]string

	for key, value := range query {
		queryParams[key] = value
	}

	return queryParams
}
