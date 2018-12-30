package config

import (
	"log"

	"github.com/spf13/viper"
)

//GetAuthToken returns the cached token from the config file
func GetAuthToken(configFile string) (token string) {
	confFiles := getConfigFile(configFile)
	for _, file := range confFiles {
		viper.AddConfigPath(file)
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error reading the config file,Error: %+v", err)
		return
	}

	token = viper.GetString(authTokenKey)
	return
}

//GetUsernamePassword return the cached username and password from the config files
func GetUsernamePassword(configFile string) (username, password string) {
	confFiles := getConfigFile(configFile)
	for _, file := range confFiles {
		viper.AddConfigPath(file)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error reading in the config file,Error: %+v", err)
		return
	}

	username = viper.GetString(usernameKey)
	password = viper.GetString(passwordKey)
	return
}
