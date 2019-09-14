package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli"
	ip "github.com/multiverse-os/ip"
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
		Commands: []cli.Command{
			cli.Command{
				Name:    "lookup",
				Aliases: []string{"l"},
				Usage:   "look up information for a given ip address",
				Flags: []cli.Flag{
					cli.Flag{
						Name:  "ip",
						Usage: "address to lookup",
						Value: "8.8.8.8",
					},
				},
			},
			cli.Command{
				Name:    "draw",
				Aliases: []string{"d"},
				Usage:   "render line on globe showing connection",
				Flags: []cli.Flag{
					cli.Flag{
						Name:  "ip",
						Usage: "address to lookup",
						Value: "8.8.8.8",
					},
				},
			},
		},
		DefaultAction: func(context *cli.Context) error {
			fmt.Println("Drawing connection to 8.8.8.8")

			ip.DrawConnection("8.8.8.8")
			return nil

		},
	})

	cmd.Run(os.Args)
}
