package goLF

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"time"
)

func NewCLI(golf *GoLF) *CLI {
	commandMap := make(map[string]*Command)
	return &CLI{commands: commandMap, logger: golf.Logger, golf: golf}
}

func (cli *CLI) AddCommand(cmd *Command) {
	flagValMap := make(map[string]*string)
	flagTypeMap := make(map[string]string)
	cli.commands[cmd.Name] = cmd
	cli.commands[cmd.Name].flagValMap = flagValMap
	cli.commands[cmd.Name].flagTypeMap = flagTypeMap
	cli.commands[cmd.Name].flags = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
}

func (cli *CLI) printUsage() {
	fmt.Printf("Usage: %s <command> [options]\n", cli.ToolName)
	fmt.Println("Available commands:")
	for _, cmd := range cli.commands {
		fmt.Printf("What is my command: %s\n", cmd.Name)
		fmt.Printf("What I do: %s\n", cmd.Description)
		fmt.Println("What I accept:")
		cmd.flags.PrintDefaults()
	}
}

func (cli *CLI) Run() {
	var ctx Context
	ctx.GoLF = cli.golf

	if len(os.Args) <= 1 || os.Args[1] == "-h" {
		cli.printUsage()
		os.Exit(1)
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		fmt.Printf("version: %s", cli.Version)
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmd, ok := cli.commands[cmdName]
	if ok {
		if err := cmd.flags.Parse(os.Args[2:]); err != nil {
			cli.logger.Errorf("Error parsing flags for command '%s': %v", cmd.Name, err)
			os.Exit(1)
		}

		flagMap := make(map[string]string)
		for flagName, flagValue := range cmd.flagValMap {
			flagType := cmd.flagTypeMap[flagName]
			if *flagValue=="no-default" {
				cli.logger.Errorf("No value provided for: %s", flagName)
				os.Exit(1)
			}
			switch flagType {
			case STRING:
				fmt.Println("String value:", flagValue)
			case INT:
				if _, err := strconv.Atoi(*flagValue); err != nil {
					cli.logger.Errorf("Cannot convert to integer: %v", err)
					os.Exit(1)
				}
			case BOOL:
				if _, err := strconv.ParseBool(*flagValue); err != nil {
					cli.logger.Errorf("Cannot convert to bool: %v", err)
					os.Exit(1)
				}
			case INT64:
				if _, err := strconv.ParseInt(*flagValue, 10, 64); err != nil {
					cli.logger.Errorf("Cannot convert to int64: %v", err)
					os.Exit(1)
				}
			case UINT:
				if _, err := strconv.ParseUint(*flagValue, 10, 0); err != nil {
					cli.logger.Errorf("Cannot convert to uint: %v", err)
					os.Exit(1)
				}
			case UINT64:
				if _, err := strconv.ParseUint(*flagValue, 10, 64); err != nil {
					cli.logger.Errorf("Cannot convert to uint64: %v", err)
					os.Exit(1)
				}
			case FLOAT64:
				if _, err := strconv.ParseFloat(*flagValue, 64); err != nil {
					cli.logger.Errorf("Cannot convert to float64: %v", err)
					os.Exit(1)
				}
			case DURATION:
				if _, err := time.ParseDuration(*flagValue); err != nil {
					cli.logger.Errorf("Cannot convert to duration: %v", err)
					os.Exit(1)
				}
			default:
				cli.logger.Errorf("Unknown type: %s", flagType)
				os.Exit(1)
			}
			flagMap[flagName] = *flagValue
		}
		ctx.Flags = flagMap
		err := cmd.Task(&ctx)
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
		cli.commands[command].flagTypeMap[value.Name] = value.Type
		cli.commands[command].flagValMap[value.Name] = cli.commands[command].flags.String(value.Name, value.Default, value.Help)
	}
}
