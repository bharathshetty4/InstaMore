package features

import (
	"fmt"
	"github.com/bharathshetty4/InstaMore/config"
	"log"
)

//LoginInstagram function is used by CLI aswell as API, called from cmd,api packages respectively
func LoginInstagram(username, password string) error {
	//get the auth from config and try to login
	cachedToken := config.GetAuthToken("")
	if cachedToken != "" {
		//try authorizing this token, continue if not able to login
		log.Print("Using the cached auth to login")
	}
	//caached token is not working, use the username and password passed/cached to login

	username, password = getUsernamePassword(username, password)
	if username == "" || password == "" {
		return fmt.Errorf("Username and Password both has to be passed")
	}
	//writes the latest username, password to the config file
	config.WriteUsernamePassword("", username, password)

	log.Printf("trying login to Instagram using username %s", username)

	return nil
}

//get the username password from config file if exist when the user isnt' passing
func getUsernamePassword(username, password string) (string, string) {
	if username == "" || password == "" {
		cachedUsername, cachedPassword := config.GetUsernamePassword("")
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
