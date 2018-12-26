package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func SetAppProperties(app *cli.App) {
	app.Name = "InstaMore"
	app.HelpName = "InstaMore"
	app.Description = "a CLI application to do more on Instagram"
	app.CustomAppHelpTemplate = "InstaMore "
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.CustomAppHelpTemplate = fmt.Sprintf("%s \nmore info at: https://github.com/bharathshetty4/InstaMore\n", cli.AppHelpTemplate)

	//Author related
	app.Author = "Bharath Kumar"
	app.Email = "shettybharath4@gmail.com"
}
