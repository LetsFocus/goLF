package goLF

import (
	"flag"

	"github.com/LetsFocus/goLF/logger"
)

const (
	STRING   = "string"
	INT      = "int"
	BOOL     = "bool"
	INT64    = "int64"
	UINT     = "uint"
	UINT64   = "uint64"
	FLOAT64  = "float64"
	DURATION = "duration"
)

type Flags struct {
	Name    string
	Type    string
	Default string
	Help    string
}

type Command struct {
	Name        string
	Description string
	flags       *flag.FlagSet
	flagValMap  map[string]*string
	flagTypeMap map[string]string
	Task        func(ctx Context) error
}

type CLI struct {
	ToolName string
	Version  string
	logger   *logger.CustomLogger
	ctx      *Context
	commands map[string]*Command
}
