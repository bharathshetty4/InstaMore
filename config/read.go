package config

import (
	"github.com/ahmdrz/goinsta"
	"github.com/spf13/viper"
	"log"
)

//WriteInstagramObject writes the Instagram object to the configFileDir
func ReadInstagramObject() (*goinsta.Instagram, error) {
	confFileDir := getConfigFileDir()

	return goinsta.Import(confFileDir + "auth")
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
