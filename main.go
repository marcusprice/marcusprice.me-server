package main

import (
	"fmt"
	"os"
	"strings"

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
		config.Port = "6969"
	}
}

func main() {
	config := Config{}
	config.SetDefaults()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		userAgent := strings.Split(c.GetHeader("User-Agent"), "/")[0]
		if userAgent == "curl" {
			businessCard, err := os.ReadFile("./public/businessCard")	
			if err != nil {
				// handle error
				fmt.Println(err)
			} else {
				c.Header("Content-Type", "text/html")
				_, _ = c.Writer.Write(businessCard)
			}
		} else {
				
		}
	})
	r.Run(config.Host + ":" + config.Port)
}

