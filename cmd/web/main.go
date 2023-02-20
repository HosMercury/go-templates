package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HosMercury/go-templates/pkg/config"
	"github.com/HosMercury/go-templates/pkg/handlers"
	"github.com/HosMercury/go-templates/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cache err")
	}

	app.TemplateCahe = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Listening on port %s", port))

	_ = http.ListenAndServe(port, nil)
}
