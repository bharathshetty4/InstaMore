package config

import (
	"log"
	"os"
)

//all config related const are defined here
const (
	usernameKey  = "username"
	passwordKey  = "password"
	authTokenKey = "auth"
)

//all conf files are added here which will be read by the viper
func getConfigFile(configFile string) []string {
	homeDir := os.Getenv("HOME")
	confFiles := []string{homeDir + "/.InstaMore/config"}

	if configFile != "" {
		confFiles = append(confFiles, configFile)
	}
	return confFiles
}

func createConfFilesIfNotExist(confFiles []string) (err error) {
	for _, file := range confFiles {
		f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
		if err := f.Close(); err != nil {
			log.Fatalf("error creating the config file %s, Error: %+v", file, err)
		}
	}
	return
}
