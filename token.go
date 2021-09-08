package cli

import (
	"fmt"
	"strings"
)

type token interface {
	toString() string
	validate() error
	compare(string) error
}

type tokenIterator struct {
	tokens []token
	pos    int
}

func (tknIter tokenIterator) current() token {
	if tknIter.pos >= 0 && tknIter.pos < len(tknIter.tokens) {
		return tknIter.tokens[tknIter.pos]
	}
	return nil
}

func (tknIter *tokenIterator) advance() {
	tknIter.pos++
}

type commandToken string

func (t commandToken) toString() string {
	return string(t)
}

func (t commandToken) validate() error { /*TODO*/ return nil }
func (t commandToken) compare(cmp string) error {
	if t.toString() != cmp {
		return fmt.Errorf("command '%s' expected, command '%s' recieved", cmp, t)
	}
	return nil
}

type flagToken string

func (t flagToken) toString() string {
	return string(t)
}

func (t flagToken) validate() error { /*TODO*/ return nil }

func (t flagToken) compare(cmp string) error {
	if t.toString() != cmp {
		return fmt.Errorf("flag '%s' expected, flag '%s' recieved", cmp, t)
	}
	return nil
}

func tokenify(words []string) tokenIterator {
	// TODO need to have better way
	ti := tokenIterator{}
	for _, word := range words {
		if strings.HasPrefix(word, "--") {
			ti.tokens = append(ti.tokens, flagToken(word[2:]))
		} else {
			ti.tokens = append(ti.tokens, commandToken(word))
		}
	}
	return ti
}
