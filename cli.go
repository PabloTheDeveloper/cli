package cli

// Command is a interface you may implement to have move control over what flags and subcommands are in the command
type Command interface {
	GetName() string
	GetUsage() string
	GetFlags() ([]Flag, error)
	GetCommands() ([]Command, error)
	GetExecutor() func(efs ExecutableCommand) error
}

// StandardCommand is used to contain the outline of a command
type StandardCommand struct {
	Name        string
	Usage       string
	Flags       []Flag
	SubCommands []Command
	Execute     func(ExecutableCommand) error
}

// GetName gets the name of the command
func (c StandardCommand) GetName() string {
	return c.Name
}

// GetUsage gets the usage of the command
func (c StandardCommand) GetUsage() string {
	return c.Usage
}

// GetFlags gets the flags from the command
func (c StandardCommand) GetFlags() ([]Flag, error) {
	return c.Flags, nil
}

// GetCommands gets subcommands from the command
func (c StandardCommand) GetCommands() ([]Command, error) {
	return c.SubCommands, nil
}

// GetExecutor gets the executing function which will be used in the command
func (c StandardCommand) GetExecutor() func(e ExecutableCommand) error { return c.Execute }

type flagType int

const (
	// Bool is a type to define what kind of flag the Flag struct will be
	Bool flagType = iota
)

// Flag is used to contain the outline of a flag
type Flag struct {
	Label string
	Usage string
	Type  flagType
}

// ExecuteCommand will execute each command in the path form by tokens passed into the program
func ExecuteCommand(cmd Command, ti *tokenIterator) error {
	cmdParser := commandParser{Command: cmd}
	var executableCmd ExecutableCommand
	if err := cmdParser.parse(ti, &executableCmd); err != nil {
		return err
	}
	return executableCmd.execute()
}
