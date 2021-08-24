package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DavidHODs/tsaw/config"
	"github.com/DavidHODs/tsaw/handlers"
	"github.com/DavidHODs/tsaw/render"
)

const portNumber = ":8080"


// main is the main application function
func main() {
	var app config.AppConfig
	
	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}