package config

import (
	"github.com/spf13/viper"
	"log"
)

//WriteUsernamePassword rites the cached username and password to the config file
func WriteUsernamePassword(username, password string) (err error) {
	confFile := getConfigFileDir()

	err = createConfFilesIfNotExist(confFile)
	if err != nil {
		log.Printf("Error creating the config file, %+v", err)
		return
	}
	viper.Set(usernameKey, username)
	viper.Set(passwordKey, password)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(confFile)
	err = viper.WriteConfig()
	if err != nil {
		log.Printf("error writing the config file to cache username/password, Error: %+v\n", err)
		return
	}

	return
}

//WriteUsernamePassword rites the cached username and password to the config file
func WriteAuthToken(authToken string) (err error) {
	confFile := getConfigFileDir()

	err = createConfFilesIfNotExist(confFile)
	if err != nil {
		log.Printf("Error creating the config file, %+v", err)
		return
	}
	viper.Set(authTokenKey, authToken)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(confFile)
	err = viper.WriteConfig()
	if err != nil {
		log.Printf("error writing the config file to cache username/password, Error: %+v\n", err)
		return
	}
	return
}
