package config

import (
	"log"
	"os"
)

//all config related const are defined here
const (
	usernameKey  = "username"
	passwordKey  = "password"
)

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	confFile := homeDir + "/.InstaMore"

	return confFile
}

func createConfFilesIfNotExist(confFileDir string) error {
	confFile := confFileDir + "/config.yaml"
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		os.MkdirAll(confFileDir, 0777)
		f, err := os.OpenFile(confFile, os.O_CREATE, 0666)
		if err != nil {
			log.Printf("unable create the config file at %s, %v", confFile, err.Error())
			return err
		}
		defer f.Close()
	}
	return nil
}
