package cli

import (
	"fmt"
	"strings"
)

type token string
type tokenIterator struct {
	tokens []token
	pos    int
}

func (ti tokenIterator) current() (token, bool) {
	if ti.pos >= 0 && ti.pos < len(ti.tokens) {
		return ti.tokens[ti.pos], true
	}
	return "", false
}
func (ti *tokenIterator) advance() bool {
	if ti.pos < len(ti.tokens)-1 {
		ti.pos++
		return true
	}
	return false
}
func (ti *tokenIterator) retreat() bool {
	if ti.pos > 0 {
		ti.pos--
		return true
	}
	return false
}

type flagToken token

func toFlagToken(tkn token) (flagToken, bool) {
	if strings.HasPrefix(string(tkn), "--") && len(tkn) >= 3 {
		return flagToken(tkn[1:]), true
	}
	return "", false
}

type tokenSet struct {
	ready map[token]int
	used  map[token]bool
}

// TODO do set moethod for tokenSet (check for initialization errors)
func (ts tokenSet) get(exp token) (int, error) {
	if ts.used[exp] {
		return -1, fmt.Errorf("'%s' already used", exp)
	}
	idx, ok := ts.ready[exp]
	if !ok {
		return -1, fmt.Errorf("'%s' not valid", exp)
	}
	ts.used[exp] = true
	return idx, nil
}

type CmdExp interface {
	Exp
	GetFlagsExp() (FlagsExp, error)
	GetSubCmdsExp() (SubCmdsExp, error)
}
type SubCmdsExp interface {
	Exp
	Get(string) (CmdExp, error)
}
type FlagsExp interface {
	Exp
	Get(string) (FlagExp, error)
}
type Exp interface {
	Parse(*tokenIterator) error
}
type FlagExp interface {
	Exp
}
