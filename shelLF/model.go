package shelLF

import "flag"

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
	Default interface{}
	Help    string
}

type Command struct {
	Name        string
	Description string
	Flags       *flag.FlagSet
	FlagMap     map[string]interface{}
	Task        func(flags map[string]interface{}) error
}

type CLI struct {
	toolName string
	commands map[string]*Command
}
