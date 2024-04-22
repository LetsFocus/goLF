package goLF

import (
	"fmt"
	"os"
	"strings"
)

func NewCLI(golf *GoLF) *CLI {
	commandMap := make(map[string]*Command)
	return &CLI{commands: commandMap, logger: golf.Logger, golf: golf}
}

func (cli *CLI) AddCommand(cmd *Command) {
	flagValMap := make(map[string]string)
	flagHelpMap := make(map[string]string)
	cli.commands[cmd.Name] = cmd
	cli.commands[cmd.Name].flagValMap = flagValMap
	cli.commands[cmd.Name].flagHelpMap = flagHelpMap
}

func (cli *CLI) printUsage() {
	fmt.Printf("Usage: %s <command> [options]\n", cli.ToolName)
	fmt.Println("Available commands:")
	for _, cmd := range cli.commands {
		fmt.Printf("What is my command: %s\n", cmd.Name)
		fmt.Printf("What I do: %s\n", cmd.Description)
		fmt.Println("What I accept:")
		for paramName := range cmd.flagValMap{
			fmt.Printf("	%s - %s\n", paramName, cmd.flagHelpMap[paramName])
		}

	}
}

func (cli *CLI) parseCommand(commands []string) (string, map[string]string) {
	commandName := commands[1]
	parameters := make(map[string]string)

	if len(commands)>=2 {
		for _, command := range commands[2:] {
			paramParts := strings.SplitN(command, "=", 2)
			if len(paramParts) != 2 {
				fmt.Printf("Invalid parameter format: %s\n", command)
				continue
			}
			paramName := strings.TrimLeft(paramParts[0], "-")
			parameters[paramName] = paramParts[1]
		}
	}

	return commandName, parameters
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

	commandName, parameters := cli.parseCommand(os.Args)
	cmd, ok := cli.commands[commandName]
	if ok {
		for paramName := range cmd.flagValMap {
			if _, ok := parameters[paramName]; ok {
				cmd.flagValMap[paramName] = parameters[paramName]
			}
		}
		ctx.Flags = cmd.flagValMap
		err := cmd.Task(&ctx)
		if err != nil {
			cli.logger.Errorf("Error executing command '%s': %v\n", cmd.Name, err)
		}
		return
	} else {
		cli.logger.Errorf("Error: Unknown command '%s'\n", commandName)
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
		cli.commands[command].flagValMap[value.Name] = ""
		cli.commands[command].flagHelpMap[value.Name] = value.Help
	}
}
