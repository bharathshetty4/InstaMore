package main

import (
	"os"
	"log"

	"github.com/urfave/cli"

	"github.com/bharathshetty4/InstaMore/cmd"
)

func main() {
	app := cli.NewApp()

	//set application related properties here
	cmd.SetAppProperties(app)

	//add commands and parameters to the application
	cmd.AddCommands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err.Error())
	}
}
