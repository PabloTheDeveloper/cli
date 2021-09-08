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

	rootCmd2 := StandardCommand{
		Name:  "formly",
		Usage: "a tool for building form on cli",
		Flags: []Flag{
			{
				Label: "v",
				Usage: "shows extra information",
				Type:  Bool,
			},
		},
		SubCommands: []Command{
			StandardCommand{
				Name:  "second",
				Usage: "a second subcommand",
				Flags: []Flag{
					{
						Label: "a",
						Usage: "some usage",
						Type:  Bool,
					},
				},
				Execute: func(ec ExecutableCommand) error {
					if ec.BoolFlags["a"] {
						fmt.Println("Here is extra aaaaaaa stuff")
					} else {
						fmt.Println("Normal aa level")
					}
					return nil
				},
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
	t2 := tokenify([]string{"formly", "--v", "second", "--a"})
	fmt.Println(ExecuteCommand(rootCmd2, &t2))

}
