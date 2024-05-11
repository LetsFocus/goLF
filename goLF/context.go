package goLF

import (
	"context"
	"errors"
	"net/http"
)

type Context struct {
	context.Context
	GoLF *GoLF
	Requester
	Response
	Flags map[string]string
}

type Handler func(ctx *Context)

func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r := Request{req: req}
	ctx := Context{Requester: &r, Response: Response{rw}}
	h(&ctx)
}

func (c *Context) Bind(dataType interface{}) error {
	return c.Requester.Bind(dataType)
}

func (c *Context) GetPathParam() string {
	return c.Requester.GetPathParam()
}

func (c *Context) GetHeader(key string) string {
	return c.Requester.GetHeader(key)
}

func (c *Context) GetHeaders() map[string][]string {
	return c.Requester.GetHeaders()
}

func (c *Context) GetParam(key string) string {
	return c.Requester.GetParam(key)
}

func (c *Context) GetParams() map[string]string {
	return c.Requester.GetParams()
}

func (c *Context) GetParamsArray() map[string][]string {
	return c.Requester.GetParamsArray()
}

// function to get all pairs of flag and values
func (c *Context) GetALLFlags() map[string]string {
	return c.Flags
}

// function to get value corresponding to a flag
func (c *Context) GetFlag(flagname string) (string, error) {
	value, exists := c.Flags[flagname]
	if exists {
		return value, nil
	} else {
		return "", errors.New("flag not found")
	}
}
