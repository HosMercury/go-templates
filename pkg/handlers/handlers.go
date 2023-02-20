package handlers

import (
	"net/http"

	"github.com/HosMercury/go-templates/pkg/config"
	"github.com/HosMercury/go-templates/pkg/render"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl", &TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello again "

	render.RenderTemplates(w, "about.page.tmpl", &TemplateData{})
}
