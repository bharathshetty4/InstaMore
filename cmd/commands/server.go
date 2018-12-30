package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

//Define all server related commands here
var (
	CommandStartServer = cli.Command{
		Name:        "server",
		Aliases:     []string{"s"},
		Description: "start a server to do RESTful calls",
		ArgsUsage:   "[GLOBAL ARGUMENTS]",

		Subcommands: []cli.Command{
			{
				Name:        "start", //start the server
				Description: "start a server to do RESTful calls",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "port",
						Usage: "port on which the server should run(default port is 5555)",
					},
				},
				Action: startServer,
			},
			{
				Name: "apilist", //list all the api's available
				Description: "Lists all the API's available	",
				Action:    apiList,
				ArgsUsage: "[GLOBAL ARGUMENTS]",
			},
		},
	}
)

func startServer(c *cli.Context) error {
	fmt.Println("starting the server... ", c.Args())
	return nil
}

func apiList(c *cli.Context) error {
	fmt.Println("All available API's are:", c.Args())
	return nil
}
