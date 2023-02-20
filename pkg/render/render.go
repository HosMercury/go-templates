package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/HosMercury/go-templates/pkg/config"
	"github.com/HosMercury/go-templates/pkg/handlers"
)

var app *config.AppConfig

var functions = template.FuncMap{}

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplates(w http.ResponseWriter, tmpl string, td handlers.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCahe
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("cache err")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("err")
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles((page))

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}

	}

	return myCache, nil
}
