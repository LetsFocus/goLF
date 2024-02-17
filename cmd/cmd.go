package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"time"

	"github.com/LetsFocus/goLF/logger"
)

func NewCMD() *CLI {
	commandMap := make(map[string]*Command)
	logger := logger.NewCustomLogger()
	return &CLI{commands: commandMap, logger: logger}
}

func (cli *CLI) AddCommand(cmd Command) {
	cli.commands[cmd.Name] = &cmd
}

func (cli *CLI) printUsage() {
	fmt.Printf("Usage: %s <command> [options]", cli.ToolName)
	fmt.Printf("Available commands:")
	for _, cmd := range cli.commands {
		fmt.Printf("What is my command: %s\n", cmd.Name)
		fmt.Printf("What I do: %s\n", cmd.Description)
		fmt.Printf("What I accept:")
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
			case *float32:
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
	flagsToCmd := flag.NewFlagSet(cli.commands[command].Name, flag.ExitOnError)
	flagMap := make(map[string]interface{})
	for _, value := range cmdFlags {
		switch value.Type {
		case STRING:
			flagMap[value.Name] = flagsToCmd.String(value.Name, value.Default.(string), value.Help)
		case INT:
			flagMap[value.Name] = flagsToCmd.Int(value.Name, value.Default.(int), value.Help)
		case INT64:
			flagMap[value.Name] = flagsToCmd.Int64(value.Name, value.Default.(int64), value.Help)
		case UINT:
			flagMap[value.Name] = flagsToCmd.Uint(value.Name, value.Default.(uint), value.Help)
		case UINT64:
			flagMap[value.Name] = flagsToCmd.Uint64(value.Name, value.Default.(uint64), value.Help)
		case FLOAT32:
			str := strconv.FormatFloat(float64(value.Default.(float32)), 'f', -1, 32)
			floatToStringFlag := flagsToCmd.String(value.Name, str, value.Help)
			stringValue := *floatToStringFlag
			floatValue, err := strconv.ParseFloat(stringValue, 32)
			if err != nil {
				cli.logger.Errorf("Error: Invalid float value '%v'\n", err)
				cli.printUsage()
				os.Exit(1)
			}
			float32Val := float32(floatValue)
			flagMap[value.Name] = &float32Val
		case FLOAT64:
			flagMap[value.Name] = flagsToCmd.Float64(value.Name, value.Default.(float64), value.Help)
		case BOOL:
			flagMap[value.Name] = flagsToCmd.Bool(value.Name, value.Default.(bool), value.Help)
		case DURATION:
			flagMap[value.Name] = flagsToCmd.Duration(value.Name, value.Default.(time.Duration), value.Help)
		}
	}

	cli.commands[command].FlagMap = flagMap
	cli.commands[command].Flags = flagsToCmd
}
