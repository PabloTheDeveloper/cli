package cli

type ExecutableCommand struct {
	// each flag type are here as maps
	// ex. boolFlags map[string]bool
	// ex. strFlags map[string]string

	// subcommand is the subcommand of the command
	// if nil, then it does not exist
	BoolFlags  map[string]bool
	subcommand *ExecutableCommand
	execute    func(efs ExecutableCommand) error
}

func (ec ExecutableCommand) Execute() error {
	if err := ec.execute(ec); err != nil {
		return err
	}
	if ec.subcommand == nil {
		return nil
	}
	return ec.subcommand.Execute()
}

// TODO build the executable command
// As the tokens are being processed, I will build
// the executable command one step at a time
