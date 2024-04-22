package goLF

import (
	"github.com/LetsFocus/goLF/logger"
)

type Flags struct {
	Name string
	Help string
}

type Command struct {
	Name        string
	Description string
	flagValMap  map[string]string
	flagHelpMap map[string]string
	Task        func(ctx *Context) error
}

type CLI struct {
	ToolName string
	Version  string
	golf     *GoLF
	logger   *logger.CustomLogger
	commands map[string]*Command
}
