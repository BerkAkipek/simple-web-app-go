package handlers

import (
	"net/http"

	"github.com/BerkAkipek/simple-web-app-go/pkg/config"
	"github.com/BerkAkipek/simple-web-app-go/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (R *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (R *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
