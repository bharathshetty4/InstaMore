package config

import (
	"log"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/spf13/viper"
)

//WriteUsernamePassword writes the cached username and password to the config file
func WriteUsernamePassword(username, password string) (err error) {
	confFileDir := getConfigFileDir()

	err = createConfFilesIfNotExist(confFileDir)
	if err != nil {
		log.Printf("Error creating the config file, %+v", err)
		return
	}
	viper.Set(passwordKey, password)
	viper.Set(usernameKey, username)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(confFileDir)
	err = viper.WriteConfig()
	if err != nil {
		log.Printf("error writing the config file to cache username/password, Error: %+v\n", err)
	}
	return
}

//WriteInstagramObject writes the Instagram object to the configFileDir
func WriteInstagramObject(instagramObject *goinsta.Instagram) error {
	confFileDir := getConfigFileDir()

	os.MkdirAll(confFileDir, 0777)
	return instagramObject.Export(confFileDir + "auth")
}
