package cli

import (
	"errors"
	"fmt"
)

type parser interface {
	parse(ti *tokenIterator, ec *ExecutableCommand) error
}

type commandParser struct {
	Command
	tokenSet
}

type tokenSet struct {
	ready map[token]parser
	used  map[token]bool
}

func (set tokenSet) get(tkn token) (parser, error) {
	if set.used[tkn] {
		return nil, fmt.Errorf("'%s' already used", tkn)
	}
	parser, _ := set.ready[tkn]
	set.used[tkn] = true
	return parser, nil
}

func (p *commandParser) setup() error {
	p.tokenSet = tokenSet{
		ready: map[token]parser{},
		used:  map[token]bool{},
	}
	flags, err := p.Command.GetFlags()
	if err != nil {
		return err
	}
	for _, flag := range flags {
		flagToken := flagToken(flag.Label)
		if err := flagToken.validate(); err != nil {
			return err
		}
		if _, ok := p.ready[flagToken]; ok {
			return fmt.Errorf("flag '%s' was declared twice", flagToken)
		}
		switch flag.Type {
		case Bool:
			p.ready[flagToken] = &boolFlagParser{Flag: flag}
		default:
			panic("unknown flags type used")
		}
	}
	cmds, err := p.Command.GetCommands()
	if err != nil {
		return err
	}
	for _, cmd := range cmds {
		cmdToken := commandToken(cmd.GetName())
		if err := cmdToken.validate(); err != nil {
			return err
		}
		if _, ok := p.ready[cmdToken]; ok {
			return fmt.Errorf("cmd '%s' was declared twice", cmdToken)
		}
		p.tokenSet.ready[cmdToken] = &commandParser{Command: cmd}
	}
	return nil
}

func (p *commandParser) parse(ti *tokenIterator, ec *ExecutableCommand) error {
	if err := p.setup(); err != nil {
		return err
	}
	ec.fn = p.GetExecutor()
	ec.BoolFlags = map[string]bool{}
	ti.advance()
	stop := false
	for tkn := ti.current(); tkn != nil && !stop; tkn = ti.current() {
		switch token := tkn.(type) {
		case commandToken:
			cmdParser, err := p.get(token)
			if err != nil {
				return err
			}
			if cmdParser != nil {
				ec.subcommand = &ExecutableCommand{}
				return cmdParser.parse(ti, ec.subcommand)
			}
			return errors.New("unexpected command")
		case flagToken:
			flagParser, err := p.get(token)
			if err != nil {
				return err
			}
			if flagParser == nil {
				return errors.New("no flag defined for token")
			}
			if err := flagParser.parse(ti, ec); err != nil {
				return err
			}
		default:
			panic("Non commandToken and Non flagToken stumbled across")
		}
	}
	return nil
}

type boolFlagParser struct {
	Flag
}

func (bfp *boolFlagParser) parse(ti *tokenIterator, ec *ExecutableCommand) error {
	token := ti.current() // garanteed to be non-nil
	if err := token.compare(bfp.Label); err != nil {
		return err
	}
	ec.BoolFlags[bfp.Label] = true
	ti.advance()
	return nil
}
