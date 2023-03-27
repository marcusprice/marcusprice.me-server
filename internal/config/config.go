package config

import (
	"encoding/json"
	"log"
	"os"
)

type configJson struct {
	App struct {
		Logging bool `json:"logging,omitempty"`
	} `json:"app,omitempty"`
	Server struct {
		Host      *string `json:"host,omitempty"`
		Port      *string `json:"port,omitempty"`
		StaticDir *string `json:"staticDir,omitempty"`
	} `json:"server,omitempty"`
}

type Config struct {
	// config
	Host      string
	Port      string
	StaticDir string
}

func (config *Config) SetDefaults() {
	config.readConfig()

	if config.Host == "" {
		config.Host = "0.0.0.0"
	}

	if config.Port == "" {
		config.Port = "6969"
	}

	if config.StaticDir == "" {
		config.StaticDir = "./public"
	}
}

func (config *Config) readConfig() {
	var configData configJson

	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Println("INFO: no config file present")
	}

	json.Unmarshal(configFile, &configData)

	log.Println(configData)

	if configData.Server.Host != nil {
		config.Host = *configData.Server.Host
	}

	if configData.Server.Port != nil {
		config.Port = *configData.Server.Port
	}

	if configData.Server.StaticDir != nil {
		config.StaticDir = *configData.Server.StaticDir
	}
}
