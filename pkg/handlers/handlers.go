package handlers

import (
	"net/http"

	"github.com/BerkAkipek/simple-web-app-go/pkg/config"
	"github.com/BerkAkipek/simple-web-app-go/pkg/models"
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
	remoteIP := r.RemoteAddr
	R.App.Session.Put(r.Context(), "remote-ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (R *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World!"

	remoteIP := R.App.Session.GetString(r.Context(), "remote-ip")
	stringMap["remote-ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
