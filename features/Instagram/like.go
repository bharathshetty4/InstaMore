package Instagram

import (
	"log"
)

func (insta Instagram) LikeAllPosts(userName string) error {
	log.Printf("Using %s to like all photos of %s", insta.Account.Username, userName)

	user, err := insta.Profiles.ByName(userName)
	if err != nil {
		log.Printf("Error getting the user details using username %s, error: %+v", userName, err)
		return err
	}

	itemLen := user.MediaCount
	log.Printf("user has %+v posts", itemLen)
	///////

	newUser := insta.NewUser()
	newUser.ID = user.ID
	newUser.Feed()
	err = newUser.Sync()
	if err != nil {
		log.Printf("Error new user the user details using username %s, error: %+v", userName, err)
		return err
	}
	err = newUser.Feed().Sync()
	if err != nil {
		log.Printf("Error new user sync the user details using username %s, error: %+v", userName, err)
		return err
	}
	userMedia := newUser.Feed().Items
	log.Printf("user has %+v posts", len(userMedia))

	return nil
}
