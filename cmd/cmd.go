package cmd

import (
	"flag"
	"fmt"
	"os"

	"time"

	"github.com/LetsFocus/goLF/logger"
)

func NewCMD() *CLI {
	commandMap := make(map[string]*Command)
	logger := logger.NewCustomLogger()
	return &CLI{ToolName: "myTool", Version: "0.0", commands: commandMap, logger: logger}
}

func (cli *CLI) AddCommand(cmd Command) {
	flagMap := make(map[string]interface{})
	cli.commands[cmd.Name] = &cmd
	cli.commands[cmd.Name].FlagMap = flagMap
	cli.commands[cmd.Name].Flags = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
}

func (cli *CLI) printUsage() {
	fmt.Printf("Usage: %s <command> [options]\n", cli.ToolName)
	fmt.Println("Available commands:")
	for _, cmd := range cli.commands {
		fmt.Printf("What is my command: %s\n", cmd.Name)
		fmt.Printf("What I do: %s\n", cmd.Description)
		fmt.Println("What I accept:")
		cmd.Flags.PrintDefaults()
	}
}

func (cli *CLI) Run() {
	if len(os.Args) <= 1 || os.Args[1] == "-h" {
		cli.printUsage()
		os.Exit(1)
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		cli.logger.Infof("version: %s", cli.Version)
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmd, ok := cli.commands[cmdName]
	if ok {
		if err := cmd.Flags.Parse(os.Args[2:]); err != nil {
			cli.logger.Errorf("Error parsing flags for command '%s': %v", cmd.Name, err)
			os.Exit(1)
		}

		flagMap := make(map[string]interface{})
		for flagName, flagValue := range cmd.FlagMap {
			switch value := flagValue.(type) {
			case *string:
				flagMap[flagName] = *value
			case *int:
				flagMap[flagName] = *value
			case *bool:
				flagMap[flagName] = *value
			case *int64:
				flagMap[flagName] = *value
			case *uint:
				flagMap[flagName] = *value
			case *uint64:
				flagMap[flagName] = *value
			case *float64:
				flagMap[flagName] = *value
			case *time.Duration:
				flagMap[flagName] = *value
			}
		}
		err := cmd.Task(flagMap)
		if err != nil {
			cli.logger.Errorf("Error executing command '%s': %v\n", cmd.Name, err)
		}
		return
	} else {
		cli.logger.Errorf("Error: Unknown command '%s'\n", cmdName)
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) AddFlags(command string, cmdFlags []Flags) {
	_, ok := cli.commands[command]
	if !ok {
		cli.logger.Errorf("Error: Invalid command '%s'\n", command)
		cli.printUsage()
		os.Exit(1)
	}

	for _, value := range cmdFlags {
		switch value.Type {
		case STRING:
			cli.commands[command].FlagMap[value.Name] = cli.commands[command].Flags.String(value.Name, value.Default.(string), value.Help)
		case INT:
			cli.commands[command].FlagMap[value.Name] = cli.commands[command].Flags.Int(value.Name, value.Default.(int), value.Help)
		case INT64:
			cli.commands[command].FlagMap[value.Name] = cli.commands[command].Flags.Int64(value.Name, value.Default.(int64), value.Help)
		case UINT:
			cli.commands[command].FlagMap[value.Name] = cli.commands[command].Flags.Uint(value.Name, value.Default.(uint), value.Help)
		case UINT64:
			cli.commands[command].FlagMap[value.Name] = cli.commands[command].Flags.Uint64(value.Name, value.Default.(uint64), value.Help)
		case FLOAT64:
			cli.commands[command].FlagMap[value.Name] = cli.commands[command].Flags.Float64(value.Name, value.Default.(float64), value.Help)
		case BOOL:
			cli.commands[command].FlagMap[value.Name] = cli.commands[command].Flags.Bool(value.Name, value.Default.(bool), value.Help)
		case DURATION:
			cli.commands[command].FlagMap[value.Name] = cli.commands[command].Flags.Duration(value.Name, value.Default.(time.Duration), value.Help)
		}
	}
}
