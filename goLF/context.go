package goLF

import (
	"context"
	"github.com/LetsFocus/goLF/goLF/model"
	"net/http"
)

type Context struct {
	context.Context
	model.GoLF
	Requester
	Response
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
