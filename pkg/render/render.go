package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/BerkAkipek/simple-web-app-go/pkg/config"
)

var app *config.AppConfig

func NewTemplate(appConf *config.AppConfig) {
	app = appConf
}

func RenderTemplate(w http.ResponseWriter, tmp string) {
	// create a template cache
	var tempCache map[string]*template.Template
	if app.UseCache {
		tempCache = app.TemplateCache
	} else {
		tempCache, _ = CreateTemplateCache()
	}

	// get requested template cache
	templ, ok := tempCache[tmp]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	err := templ.Execute(buf, nil)
	if err != nil {
		log.Fatal(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tempCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tempCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tempCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tempCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tempCache, err
			}
		}

		tempCache[name] = templateSet
	}
	return tempCache, nil
}
