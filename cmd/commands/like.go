package commands

import (
	"fmt"
	"log"

	"github.com/urfave/cli"

	"github.com/bharathshetty4/InstaMore/features/Instagram"
)

const (
	user = "user"
)

//Define all login related commands here
var (
	CommandLike = cli.Command{
		Name:        "like",
		Aliases:     []string{"li"},
		Usage:       "Like all or given posts of an user",
		Description: "Like all or given posts of an user",
		ArgsUsage:   "[GLOBAL ARGUMENTS]",

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  user,
				Usage: "username of the user profile, whose posts to be liked (Ex: bharath.shetty__)",
			},
		},
		Action: likePosts,
	}
)

func likePosts(c *cli.Context) error {
	instaObject, err := Instagram.InstagramNew()
	if err != nil {
		return fmt.Errorf("Unable to get the auth token, Please login using below command,\n\n" +
			"InstaMore login --username your_instagram_username --password your_instagram_password\n")
	}
	log.Printf("Successfully Logged in using username: %s", instaObject.Account.Username)
	if c.String(user) == "" {
		return fmt.Errorf("Provide the user id of the user using below command,\n\n" +
			"InstaMore like --user <profile_user_name>\n")
	}
	instaObject.LikeAllPosts(c.String(user))

	return nil
}
