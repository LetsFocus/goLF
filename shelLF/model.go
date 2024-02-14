package shelLF

type Command struct {
	Name        string
	Description string
	Task        interface{}
}

type CommandManager struct {
	commands []*Command
}
