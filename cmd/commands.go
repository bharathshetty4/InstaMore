package cmd

import (
	"github.com/urfave/cli"

	"github.com/bharathshetty4/InstaMore/cmd/commands"
)

func AddCommands(app *cli.App) {
	app.Commands = []cli.Command{
		commands.CommandLogin,
		commands.CommandStartServer,
	}
}
