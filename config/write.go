package config

import (
	"github.com/spf13/viper"
	"log"
)

//WriteUsernamePassword rites the cached username and password to the config file
func WriteUsernamePassword(configFile, username, password string) (err error) {
	confFiles := getConfigFile(configFile)

	//create conf files if not exist
	createConfFilesIfNotExist(confFiles)

	for _, file := range confFiles {
		viper.AddConfigPath(file)
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("error reading in the config file,Error: %+v", err)
		return
	}

	viper.Set(usernameKey, username)
	viper.Set(passwordKey, password)
	return
}
