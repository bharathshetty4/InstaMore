package Instagram

import (
	"log"

	"github.com/ahmdrz/goinsta"

	"github.com/bharathshetty4/InstaMore/config"
)

type Instagram struct {
	*goinsta.Instagram
}

func InstagramNew() (Instagram, error) {
	InstaObject, err := config.ReadInstagramObject()
	if err != nil {
		log.Printf("Unable to get the auth token, Please login using below command,\n" +
			"InstaMore login --username your_instagram_username --password your_instagram_password\n")
		return Instagram{}, err
	}
	err = InstaObject.Account.Sync()
	if err != nil {
		return Instagram{}, err
	}
	return Instagram{InstaObject}, nil
}
