package cli

// ExecutableCommand contains all the relevant information in order to execute a command based off flags declared
type ExecutableCommand struct {
	BoolFlags  map[string]bool
	subcommand *ExecutableCommand
	fn         func(efs ExecutableCommand) error
}

func (ec ExecutableCommand) execute() error {
	if err := ec.fn(ec); err != nil {
		return err
	}
	if ec.subcommand == nil {
		return nil
	}
	return ec.subcommand.execute()
}
