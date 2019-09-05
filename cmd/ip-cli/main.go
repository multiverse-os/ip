package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli"
)

func main() {
	// TODO: Ideally will just merge in or basically forward all existing `ip
	// commands` and use the command name `ip`, so basically it will function as a
	// way of adding in more consistent functionality (like adding JSON output to
	// every command, and providing more functionality, while keeping all the
	// original functionality and expected usage)

	cmd := cli.New(&cli.CLI{
		Name:    "ip-cli",
		Version: cli.Version{Major: 0, Minor: 1, Patch: 0},
		Usage:   "Specify a command, and one ip address or more",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("No actions have been defined yet")
			return nil
		},
	})

	cmd.Run(os.Args)
}
