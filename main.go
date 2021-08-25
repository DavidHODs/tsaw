package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DavidHODs/tsaw/config"
	"github.com/DavidHODs/tsaw/handlers"
	"github.com/DavidHODs/tsaw/render"
	"github.com/DavidHODs/tsaw/routes"

	"github.com/alexedwards/scs"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager


// main is the main application function
func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session
	
	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Listening on port %s", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes.Routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}