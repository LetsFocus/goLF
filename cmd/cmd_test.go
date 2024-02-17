package cmd

import (
	"flag"
	"os"
	"testing"
	"time"
)

func TestNewCMD(t *testing.T) {
	testcases := []struct {
		desc        string
		toolName    string
		toolVersion string
	}{
		{
			desc:        "create CLI instance with provided tool name",
			toolName:    "myCLI",
			toolVersion: "0.1",
		},
	}

	for i, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			cli := NewCMD()

			if cli == nil {
				t.Errorf("Test[%d] FAILED: NewCMD(%s) returned nil, expected non-nil CLI instance", i, tc.toolName)
			}

			if cli != nil && len(cli.commands) != 0 {
				t.Errorf("Test[%d] FAILED: Expected CLI commands map to be empty, got commands: %v", i, cli.commands)
			}
		})
	}
}

func TestAddCommand(t *testing.T) {
	cli := NewCMD()

	testCommand := Command{
		Name:        "testCommand",
		Description: "This is a test command",
		Flags:       flag.NewFlagSet("testCommandFlags", flag.ExitOnError),
		FlagMap:     make(map[string]interface{}),
		Task: func(flags map[string]interface{}) error {
			return nil
		},
	}

	cli.AddCommand(testCommand)

	_, exists := cli.commands[testCommand.Name]

	if !exists {
		t.Errorf("Failed to add command %s to CLI", testCommand.Name)
	}
}

func TestAddFlags(t *testing.T) {
	cli := NewCMD()

	testCommand := Command{
		Name:        "testCommand",
		Description: "This is a test command",
	}

	cli.AddCommand(testCommand)

	testFlags := []Flags{
		{Name: "stringFlag", Type: STRING, Default: "defaultStringValue", Help: "String flag"},
		{Name: "intFlag", Type: INT, Default: 42, Help: "Integer flag"},
		{Name: "int64Flag", Type: INT64, Default: int64(42), Help: "Int64 flag"},
		{Name: "uintFlag", Type: UINT, Default: uint(42), Help: "Uint flag"},
		{Name: "uint64Flag", Type: UINT64, Default: uint64(42), Help: "Uint64 flag"},
		{Name: "float64Flag", Type: FLOAT64, Default: 42.0, Help: "Float64 flag"},
		{Name: "float32Flag", Type: FLOAT32, Default: float32(42.0), Help: "Float32 flag"},
		{Name: "boolFlag", Type: BOOL, Default: true, Help: "Boolean flag"},
		{Name: "durationFlag", Type: DURATION, Default: 5 * time.Second, Help: "Duration flag"},
	}

	cli.AddFlags("testCommand", testFlags)

	addedCommand, ok := cli.commands["testCommand"]
	if !ok {
		t.Fatal("Failed to retrieve added command")
	}
	if addedCommand.Flags == nil {
		t.Fatal("Flags were not initialized for the command")
	}

	for _, flag := range testFlags {
		_, ok := addedCommand.FlagMap[flag.Name]
		if !ok {
			t.Errorf("Flag '%s' was not added to the command", flag.Name)
		}
	}
}

func TestCLI_Run(t *testing.T) {
	cli := NewCMD()

	testCommand := Command{
		Name:        "testCommand",
		Description: "This is a test command",
		Task: func(flags map[string]interface{}) error {
			return nil
		},
	}
	cli.AddCommand(testCommand)
	testFlags := []Flags{
		{Name: "stringFlag", Type: STRING, Default: "defaultStringValue", Help: "String flag"},
		{Name: "intFlag", Type: INT, Default: 42, Help: "Integer flag"},
		{Name: "int64Flag", Type: INT64, Default: int64(42), Help: "Int64 flag"},
		{Name: "uintFlag", Type: UINT, Default: uint(42), Help: "Uint flag"},
		{Name: "uint64Flag", Type: UINT64, Default: uint64(42), Help: "Uint64 flag"},
		{Name: "float64Flag", Type: FLOAT64, Default: 42.0, Help: "Float64 flag"},
		{Name: "float32Flag", Type: FLOAT32, Default: float32(42.0), Help: "Float32 flag"},
		{Name: "boolFlag", Type: BOOL, Default: true, Help: "Boolean flag"},
		{Name: "durationFlag", Type: DURATION, Default: 5 * time.Second, Help: "Duration flag"},
	}

	cli.AddFlags("testCommand", testFlags)

	os.Args = []string{"myCLI", "testCommand", "-stringFlag=test", "-intFlag=42", "-boolFlag=true"}
	cli.Run()

	os.Args = []string{"myCLI", "testCommand", "-stringFlag=test", "-intFlag=42", "-boolFlag=true", "-int64Flag=420", "-uintFlag=42", "-uint64Flag=420", "-float64Flag=4.2", "-float32Flag=4.2", "-durationFlag=5s"}
	cli.Run()
}
