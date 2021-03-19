package main

import (
	"fmt"
	"github.com/polyhistor/go-service/pkg/config"
	"github.com/polyhistor/go-service/pkg/handlers"
	"github.com/polyhistor/go-service/pkg/render"
	"log"
	"net/http"
)

var portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %S", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
