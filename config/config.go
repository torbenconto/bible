package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	ApiKey string `yaml:"api_key"`
}

func InitApiKey(key string) {
	// Get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Check if config file exists
	configPath := home + "/.bible/config.yaml"
	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create the config file
			_, err = os.Create(configPath)
			if err != nil {
				log.Fatal(err)
			}

			// Write the API key to the config file
			err = os.WriteFile(configPath, []byte("api_key: "+key), 0644)
			if err != nil {
				log.Fatal(err)
			}

			return
		}
	}
}

func GetApiKey() string {
	// Get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Check if config file exists
	configPath := home + "/.bible/config.yaml"
	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
	}

	// Read the config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the API key
	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config.ApiKey
}
