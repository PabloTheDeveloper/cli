package cli

type Command interface {
	GetName() string
	GetUsage() string
	GetFlags() ([]Flag, error)
	GetCommands() ([]Command, error)
	GetExecutor() func(efs ExecutableCommand) error
}

type StandardCommand struct {
	Name        string
	Usage       string
	Flags       []Flag
	SubCommands []Command
	Execute     func(ExecutableCommand) error
}

func (c StandardCommand) GetName() string {
	return c.Name
}

func (c StandardCommand) GetUsage() string {
	return c.Usage
}

func (c StandardCommand) GetFlags() ([]Flag, error) {
	return c.Flags, nil
}
func (c StandardCommand) GetCommands() ([]Command, error) {
	return c.SubCommands, nil
}
func (c StandardCommand) GetExecutor() func(e ExecutableCommand) error {
	return c.Execute
}

type flagType int

const (
	Bool flagType = iota
)

type Flag struct {
	Label string
	Usage string
	Type  flagType
}

func ExecuteCommand(cmd Command, ti *tokenIterator) error {
	cmdParser := commandParser{Command: cmd}
	var executableCmd ExecutableCommand
	if err := cmdParser.parse(ti, &executableCmd); err != nil {
		return err
	}
	return executableCmd.Execute()
}
