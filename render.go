package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template: ", err)
		return
	}
}
