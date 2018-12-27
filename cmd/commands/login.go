package commands

import (
	"fmt"

	"github.com/urfave/cli"
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
	//get the auth from config and try to login

	//success: return, else get username and password from config

	if c.String(username) == "" || c.String(password) == "" {
		return fmt.Errorf("Username and Password both has to be passed")
	}

	fmt.Printf("Using the username %s to login\n\n", c.String(username))

	return nil
}
