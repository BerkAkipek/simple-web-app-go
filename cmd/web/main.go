package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BerkAkipek/simple-web-app-go/pkg/config"
	"github.com/BerkAkipek/simple-web-app-go/pkg/handlers"
	"github.com/BerkAkipek/simple-web-app-go/pkg/render"
)

const portNumber string = ":8080"

func main() {
	var app config.AppConfig

	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tempCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Server listening on http://localhost%v/\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
