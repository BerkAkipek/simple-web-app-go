package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BerkAkipek/simple-web-app-go/pkg/config"
	"github.com/BerkAkipek/simple-web-app-go/pkg/handlers"
	"github.com/BerkAkipek/simple-web-app-go/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// Change This to true iin Production
	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	app.Session = session

	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tempCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	fmt.Printf("Server listening on http://localhost%v/\n", portNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	server.ListenAndServe()
	log.Fatal(err)
}
