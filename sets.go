package cli

import (
	"fmt"
)

func (cmd Command) GetFlagsExp() (FlagsExp, error) { //TODO SETUP map for quicker retrieval (use ExpSet)
	return cmd.flags, nil
}
func (cmd Command) GetSubCmdsExp() (SubCmdsExp, error) { //TODO SETUP map for quicker retrieval (use ExpSet)
	return cmd.scmds, nil
}

// FlagSet satisfies FlagsExp
type FlagSet struct {
	Flags []FlagExp // used to list as a struct
	tokenSet
}

func (fs FlagSet) Get(tkn token) (FlagExp, error) {
	flagToken, ok := toFlagToken(tkn)
	if !ok {
		return nil, nil
	}
	idx, err := fs.tokenSet.get(token(flagToken))
	if err != nil {
		return nil, fmt.Errorf("flag: %w", err)
	}
	return fs.Flags[idx], nil
}

// SubCmdSet satisfies SubCmdsExp
type SubCmdSet struct {
	SubCommands []CmdExp // used to list as a struct
	IsRequired  bool
	tokenSet
}

func (scs SubCmdSet) Get(tkn token) (CmdExp, error) {
	idx, err := scs.tokenSet.get(tkn)
	if err != nil {
		return nil, fmt.Errorf("subcommand: %w", err)
	}
	return scs.SubCommands[idx], nil
}
