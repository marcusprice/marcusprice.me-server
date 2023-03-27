package config

import (
	"encoding/json"
	"log"
	"os"
)

type configJson struct {
	App struct {
		Logging *bool `json:"logging,omitempty"`
	} `json:"app,omitempty"`
	Server struct {
		Host      *string `json:"host,omitempty"`
		Port      *string `json:"port,omitempty"`
		StaticDir *string `json:"staticDir,omitempty"`
	} `json:"server,omitempty"`
}

type Config struct {
	Logging   bool
	Host      string
	Port      string
	StaticDir string
}

func (config *Config) Initialize() {
	config.setDefaults()
	config.readConfig()
}

func (config *Config) readConfig() {
	var configData configJson

	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Println("INFO: no config file present")
	}
	json.Unmarshal(configFile, &configData)

	if configData.Server.Host != nil {
		config.Host = *configData.Server.Host
	}

	if configData.Server.Port != nil {
		config.Port = *configData.Server.Port
	}

	if configData.Server.StaticDir != nil {
		config.StaticDir = *configData.Server.StaticDir
	}

	if configData.App.Logging != nil {
		config.Logging = *configData.App.Logging
	}
}

func (config *Config) setDefaults() {
	config.Logging = false
	config.Host = "0.0.0.0"
	config.Port = "6969"
	config.StaticDir = "./public"
}
