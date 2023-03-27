package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marcusprice/marcusprice.me-server/internal/config"
)

func main() {
	config := config.Config{}
	config.Initialize()

	r := chi.NewRouter()
	if config.Logging {
		r.Use(middleware.Logger)
	}

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

	serverAddress := config.Host + ":" + config.Port
	log.Println("starting server at " + serverAddress)
	err := http.ListenAndServe(serverAddress, r)
	if err != nil {
		log.Fatal(err)
	}
}
