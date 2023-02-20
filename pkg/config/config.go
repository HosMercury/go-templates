package config

import (
	"log"
	"text/template"
)

type AppConfig struct {
	UseCache     bool
	TemplateCahe map[string]*template.Template
	InfoLog      *log.Logger
}
