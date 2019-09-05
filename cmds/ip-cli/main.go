package main

import (
	"fmt"

	cli "github.com/multiverse-os/cli-framework"
)

func main() {
	// NOTE: This makes more sense as 'cmd' over 'app', because the application
	// version (the backing library or protocol) is separate from the CLI version.
	cmd := cli.New(&cli.CLI{
		Name:    "Example",
		Version: cli.Version{Major: 0, Minor: 1, Patch: 1},
		Usage:   "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("Example output in response to a command (action)")
			return nil
		},
	})

	cmd.Run(os.Args)
}
