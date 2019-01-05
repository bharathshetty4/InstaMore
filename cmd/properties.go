package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/bharathshetty4/InstaMore/cmd/commands"
)

//SetAppProperties adds CLI properties
func SetAppProperties(app *cli.App) {
	app.Name = "InstaMore"
	app.HelpName = "InstaMore"
	app.Description = "CLI application to do more on Instagram"
	app.Usage = "CLI application to do more on Instagram"
	app.CustomAppHelpTemplate = "InstaMore "
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.CustomAppHelpTemplate = fmt.Sprintf("%s \nmore info at: https://github.com/bharathshetty4/InstaMore\n", cli.AppHelpTemplate)

	//Author related
	app.Author = "Bharath Kumar"
	app.Email = "shettybharath4@gmail.com"
}

//AddCommands add commands list to the CLI commands
func AddCommands(app *cli.App) {
	app.Commands = []cli.Command{
		commands.CommandLogin,
		commands.CommandLike,
		commands.CommandStartServer,
	}
}
