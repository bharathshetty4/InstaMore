package cmd

import (
	//"fmt"
	"github.com/urfave/cli"
)

//Define all commands here and keep it with us, we can use it whenever it is needed

var (
	CommandStartServer = cli.Command{
		Name:        "server",
		Aliases:     []string{"s"},
		Description: "start a server to do RESTful calls",
		//Usage: "[OPTIONS] [GLOBAL ARGUMENTS]",
		ArgsUsage: "[GLOBAL ARGUMENTS]",
		//UsageText: "[OPTIONS] [GLOBAL ARGUMENTS]",

		Subcommands: []cli.Command{
			{
				Name: "start", //list all the api's available
				Description: "start a server to do RESTful calls",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "port",
						Usage: "port on which the server should run",
					},
				},
			},
			{
				Name: "apilist", //list all the api's available
				Description: "Lists all the API's available	",
			},
		},
		//Action: func(c *cli.Context) error {
		//	fmt.Println("completed task: ", c.Args().First())
		//	return nil
		//},
	}
)

func AddCommands(app *cli.App) {
	app.Commands = []cli.Command{
		CommandStartServer,
	}
}
