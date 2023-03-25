package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Host      string
	Port      string
	StaticDir string
}

func (config *Config) SetDefaults() {
	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Port == "" {
		config.Port = "6969"
	}

	if config.StaticDir == "" {
		config.StaticDir = "./public"
	}
}

func main() {
	config := Config{}
	config.SetDefaults()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// serve static files at root dir
	fs := http.FileServer(http.Dir(config.StaticDir))
	r.Handle("/*", http.StripPrefix("/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		if strings.Contains(userAgent, "curl") {
			http.ServeFile(w, r, "./assets/businessCard")
		} else {
			http.ServeFile(w, r, config.StaticDir+"/index.html")
		}
	})

	log.Println("Starting on port: " + config.Port)
	err := http.ListenAndServe("0.0.0.0:"+config.Port, r)
	if err != nil {
		log.Fatal(err)
	}
}
