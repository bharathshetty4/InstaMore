package config

import (
	"log"

	"github.com/spf13/viper"
)

//GetAuthToken returns the cached token from the config file
func GetAuthToken() (token string) {
	confFile := getConfigFileDir()
	viper.AddConfigPath(confFile)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error reading the auth token from config file, Error: %+v", err)
		return
	}

	token = viper.GetString(authTokenKey)
	return
}

//GetUsernamePassword return the cached username and password from the config files
func GetUsernamePassword() (username, password string) {
	confFile := getConfigFileDir()
	viper.AddConfigPath(confFile)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error reading the username/password from config file, Error: %+v", err)
		return
	}

	username = viper.GetString(usernameKey)
	password = viper.GetString(passwordKey)
	return
}
