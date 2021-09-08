package cli

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	rootCmd := StandardCommand{
		Name:  "formly",
		Usage: "a tool for building form on cli",
		Flags: []Flag{
			{
				Label: "v",
				Usage: "shows extra information",
				Type:  Bool,
			},
		},
		Execute: func(ec ExecutableCommand) error {
			if ec.BoolFlags["v"] {
				fmt.Println("Here is extra verbose stuff")
			} else {
				fmt.Println("Normal verbose level")
			}
			return nil
		},
	}

	tokens := tokenify([]string{"formly", "--v"})
	fmt.Println(ExecuteCommand(rootCmd, &tokens))
}
