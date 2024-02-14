package shelLF

import (
	"flag"
	"log"
	"reflect"
	"strings"
)

func NewCommandManager() *CommandManager {
	return &CommandManager{}
}

func (cli *CommandManager) AddCommand(name string, descrition string, task interface{}) {
	command := &Command{
		Name:        name,
		Description: descrition,
		Task:        task,
	}
	cli.commands = append(cli.commands, command)
}

func (cli *CommandManager) Run() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 || args[0] == "help" {
		log.Printf("Usage: go run main.go <command> [arguments]\n")
		log.Printf("Available commands:\n")
		for _, cmd := range cli.commands {
			log.Printf("%s: %s\n", cmd.Name, cmd.Description)
		}
		return
	}

	commandName := args[0]
	for _, cmd := range cli.commands {
		if cmd.Name == commandName {
			var cmdArgs []reflect.Value
			numCmdArgs := reflect.TypeOf(cmd.Task).NumIn()
			numArgs := len(args[1:])
			if numArgs != numCmdArgs {
				log.Printf("Provided arguments didn't match expected arguments of %s", cmd.Name)
				return
			}
			for _, arg := range args[1:] {
				parts := strings.SplitN(arg, "=", 2)
				if len(parts) != 2 {
					log.Printf("Invalid argument format: %s", arg)
					return
				}
				cmdArgs = append(cmdArgs, reflect.ValueOf(parts[1]))
			}
			reflect.ValueOf(cmd.Task).Call(cmdArgs)
			return
		}
	}

	log.Printf("Command not found: %s\n", commandName)
}
