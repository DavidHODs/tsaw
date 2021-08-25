package routes

import (
	"net/http"

	"github.com/DavidHODs/tsaw/config"
	"github.com/DavidHODs/tsaw/handlers"
	"github.com/DavidHODs/tsaw/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	chi := chi.NewRouter()

	chi.Use(middleware.Recoverer)
	chi.Use(middlewares.WriteToConsole)
	chi.Use(middlewares.NoSurf)

	chi.Get("/", handlers.Repo.Home)
	chi.Get("/about", handlers.Repo.About)

	return chi
}