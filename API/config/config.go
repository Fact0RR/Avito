package config

import (
	"encoding/json"
	"log"
	"os"

)

type Config struct {
	Port        				string `json:"port"`
	DataBaseString				string `json:"dataBaseString"`
	TokenAdmin					string `json:"token_admin"`
	TokenUser					string `json:"token_user"`
	LogLevel					string `json:"log_level"`
	LoggingHandler				bool   `json:"log_handlers"`
}

func GetConfig() *Config {

	content, err := os.ReadFile("API/config/config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &config
}