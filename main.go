package main

import (
	"github.com/bharathshetty4/InstaMore/cmd"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	//set application related properties here
	cmd.SetAppProperties(app)

	//add commands and parameters to the application
	cmd.AddCommands(app)


	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
