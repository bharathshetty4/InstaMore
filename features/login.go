package features

import (
	"fmt"
	"github.com/ahmdrz/goinsta"
	"log"

	"github.com/bharathshetty4/InstaMore/config"
)

//LoginInstagram function is used by CLI aswell as API, called from cmd,api packages respectively
func LoginInstagram(username, password string) error {

	cachedUsername, cachedPassword := config.GetUsernamePassword()
	username, password = getUsernamePassword(username, password, cachedUsername, cachedPassword)
	if username == "" || password == "" {
		return fmt.Errorf("run this command to login to Instagram,\n" +
			"InstaMore login --username your_instagram_username --password your_instagram_password\n")
	}

	//get the auth from config and try to login
	InstaObject, err := config.ReadInstagramObject()
	if err == nil && InstaObject.Account.Username == username {
		//try authorizing this token, continue if not able to login
		log.Printf("Already signed-in using: %s", InstaObject.Account.Username)
		return nil
	}

	//writes the latest username, password to the config file
	config.WriteUsernamePassword(username, password)

	//try logging in to instagram
	log.Printf("trying login to Instagram using username: %s", username)
	err = loginAndSaveSession(username, password)
	if err != nil {
		log.Printf("Please check the username and password you are using, Retry with Below format\n" +
			"InstaMore login --username your_instagram_username --password your_instagram_password\n")
		return err
	}

	return nil
}

//get the username password from config file if exist when the user isnt' passing
func getUsernamePassword(username, password, cachedUsername, cachedPassword string) (string, string) {
	if username == "" || password == "" {
		if username == "" && cachedUsername != "" {
			username = cachedUsername
			log.Printf("Using the cached username %s for login", cachedUsername)
		}

		if password == "" && cachedPassword != "" {
			password = cachedPassword
			log.Printf("Using the cached password for login")
		}
	}
	return username, password
}

func loginAndSaveSession(username, password string) error {
	instaStruct := goinsta.New(username, password)
	err := instaStruct.Login()
	if err != nil {
		log.Printf("Unable to login with the given credentials (%s/******), error: %+v", username, err)
		return err
	}
	err = config.WriteInstagramObject(instaStruct)
	if err != nil {
		log.Fatalf("Unable to cache the login session, Please login again, error: %+v", err)
		return err
	}
	return nil
}
