package goLF

import (
	"context"
	"github.com/LetsFocus/goLF/goLF/model"
)

type Context struct {
	context.Context
	model.GoLF
	Request
	Response
}
