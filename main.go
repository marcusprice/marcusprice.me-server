package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Host string
	Port string
}

func (config *Config) SetDefaults() {
	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Port == "" {
		config.Port = "6669"
	}
}

func main() {
	config := Config{}
	config.SetDefaults()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(config.Host + ":" + "69420")
}
