package commands

import (
	"github.com/urfave/cli"

	"github.com/bharathshetty4/InstaMore/features"
)

const (
	username = "username"
	password = "password"
)

//Define all login related commands here
var (
	CommandLogin = cli.Command{
		Name:        "login",
		Aliases:     []string{"l"},
		Usage:       "login to your instagram account by providing username and password",
		Description: "login to your instagram account by providing username and password",
		ArgsUsage:   "[GLOBAL ARGUMENTS]",

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  username,
				Usage: "Username of the Instagram Account",
			},
			cli.StringFlag{
				Name:  password,
				Usage: "Password of the Instagram Account",
			},
		},
		Action: loginInstagram,
	}
)

func loginInstagram(c *cli.Context) error {
	features.LoginInstagram(c.String(username), c.String(password))
	return nil
}
