package goLF

import (
	"flag"
	"testing"

	"github.com/LetsFocus/goLF/logger"
)

func TestNewCLI(t *testing.T) {
	golf := &GoLF{Logger: logger.NewCustomLogger()}
	cli := NewCLI(golf)

	if cli == nil {
		t.Error("NewCLI should return a non-nil CLI object")
	}

	if cli!=nil && len(cli.commands) != 0 {
		t.Error("Commands map should be initialized and empty")
	}

	if cli!=nil && cli.logger != golf.Logger {
		t.Error("Logger field not set correctly")
	}

	if cli!=nil && cli.golf != golf {
		t.Error("GoLF field not set correctly")
	}
}

func TestAddCommand(t *testing.T) {
	cli := &CLI{
		ToolName: "test_tool",
		commands: make(map[string]*Command),
		logger:   logger.NewCustomLogger(),
		golf:     &GoLF{Logger: logger.NewCustomLogger()},
	}

	cmd := &Command{Name: "testCmd", Description: "Test Command"}
	cli.AddCommand(cmd)

	if len(cli.commands) != 1 {
		t.Error("Command not added to the commands map")
	}

	addedCmd := cli.commands["testCmd"]
	if addedCmd == nil {
		t.Error("Command not found in the commands map")
	}

	if addedCmd!=nil && len(addedCmd.flagValMap) != 0 || len(addedCmd.flagTypeMap) != 0 {
		t.Error("Flag maps should be initialized and empty")
	}

	if addedCmd.flags == nil {
		t.Error("FlagSet should be created")
	}
}

// Test AddFlags function
func TestAddFlags(t *testing.T) {
	cli := &CLI{
		commands: make(map[string]*Command),
		logger:   logger.NewCustomLogger(),
	}

	// Define test commands
	cmd1 := &Command{Name: "cmd1", Description: "Command 1"}
	cmd1.flags = flag.NewFlagSet("cmd1", flag.ExitOnError)
	cli.AddCommand(cmd1)

	cmd2 := &Command{Name: "cmd2", Description: "Command 2"}
	cmd2.flags = flag.NewFlagSet("cmd2", flag.ExitOnError)
	cli.AddCommand(cmd2)

	testCases := []struct {
		command  string
		cmdFlags []Flags
	}{
		{
			command: "cmd1",
			cmdFlags: []Flags{
				{Name: "flag1", Type: STRING, Default: "", Help: "Flag 1 description"},
				{Name: "flag2", Type: INT, Default: "0", Help: "Flag 2 description"},
				{Name: "flag3", Type: BOOL, Default: "false", Help: "Flag 3 description"},
				{Name: "flag4", Type: INT64, Default: "0", Help: "Flag 4 description"},
				{Name: "flag5", Type: UINT, Default: "0", Help: "Flag 5 description"},
				{Name: "flag6", Type: UINT64, Default: "0", Help: "Flag 6 description"},
				{Name: "flag7", Type: FLOAT64, Default: "0.0", Help: "Flag 7 description"},
				{Name: "flag8", Type: DURATION, Default: "0", Help: "Flag 8 description"},
			},
		},
		{
			command: "cmd2",
			cmdFlags: []Flags{
				{Name: "flag9", Type: STRING, Default: "", Help: "Flag 9 description"},
				{Name: "flag10", Type: INT, Default: "0", Help: "Flag 10 description"},
				{Name: "flag11", Type: BOOL, Default: "false", Help: "Flag 11 description"},
				{Name: "flag12", Type: INT64, Default: "0", Help: "Flag 12 description"},
				{Name: "flag13", Type: UINT, Default: "0", Help: "Flag 13 description"},
				{Name: "flag14", Type: UINT64, Default: "0", Help: "Flag 14 description"},
				{Name: "flag15", Type: FLOAT64, Default: "0.0", Help: "Flag 15 description"},
				{Name: "flag16", Type: DURATION, Default: "0", Help: "Flag 16 description"},
			},
		},
	}

	for _, tc := range testCases {
		cli.AddFlags(tc.command, tc.cmdFlags)

		cmd, ok := cli.commands[tc.command]
		if !ok {
			t.Errorf("Command '%s' not found", tc.command)
		}

		for _, flag := range tc.cmdFlags {
			if cmd.flagTypeMap[flag.Name] != flag.Type {
				t.Errorf("Expected flagTypeMap[%s] to be '%s', got '%s'", flag.Name, flag.Type, cmd.flagTypeMap[flag.Name])
			}

			if _, ok := cmd.flagValMap[flag.Name]; !ok {
				t.Errorf("Flag '%s' not found in flagValMap", flag.Name)
			}
		}
	}
}