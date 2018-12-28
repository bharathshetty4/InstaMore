package features

import (
	"fmt"
)
//LoginInstagram function is used by CLI aswell as API, called from cmd,api packages respectively
func LoginInstagram(username, password string) error {

	//get the auth from config and try to login

	//success: return, else get username and password from config

	if username == "" || password == "" {
		return fmt.Errorf("Username and Password both has to be passed")
	}

	fmt.Printf("Using the username %s to login\n\n", username)

	return nil
}
