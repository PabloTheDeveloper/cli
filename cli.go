package cli

import (
	"errors"
	"fmt"
)

// Command satisfies CmdExp
type Command struct {
	Flags       FlagSet
	SubCommands SubCmdSet
	flags       FlagsExp
	scmds       SubCmdsExp
}

func (cmd Command) Parse(ti *tokenIterator) error {
	flags, err := cmd.GetFlagsExp()
	if err != nil {
		return err
	}
	scmds, err := cmd.GetSubCmdsExp()
	if err != nil {
		return err
	}

	if err := flags.Parse(ti); err != nil {
		return err
	}
	return scmds.Parse(ti)
}

func (fs FlagSet) Parse(ti *tokenIterator) error {
	for token, ok := ti.current(); ok; token, ok = ti.current() {
		if flag, err := fs.Get(token); err != nil {
			return err
		} else if flag == nil {
			break
		} else if err := flag.Parse(ti); err != nil {
			return err
		}
	}
	return nil
}

func (scs SubCmdSet) Parse(ti *tokenIterator) error {
	token, ok := ti.current()
	if scs.IsRequired && !ok {
		return errors.New("*A subcommand is required for command*")
	}
	cmd, err := scs.Get(token)
	if err != nil {
		return err
	}
	if cmd == nil {
		return nil
	}
	return cmd.Parse(ti)
}

type BoolFlag struct {
	name  flagToken
	value bool
}

func (bf BoolFlag) Parse(ti *tokenIterator) error {
	token, ok := ti.current()
	if !ok {
		return errors.New("no more tokens")
	}
	flagToken, ok := toFlagToken(token)
	if !ok {
		return errors.New("Not a flag")
	}
	if flagToken != bf.name {
		return fmt.Errorf("expected %s but recieved %s", bf.name, flagToken)
	}

	ti.advance()
	return nil
}
