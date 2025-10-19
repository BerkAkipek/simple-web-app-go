package main

import (
	"net/http"
	"path/filepath"

	"github.com/BerkAkipek/simple-web-app-go/pkg/config"
	"github.com/BerkAkipek/simple-web-app-go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	staticDir := http.Dir(filepath.Join(".", "static"))
	fileServer := http.FileServer(staticDir)
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
