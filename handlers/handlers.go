package handlers

import (
	"net/http"

	"github.com/DavidHODs/tsaw/config"
	"github.com/DavidHODs/tsaw/models"
	"github.com/DavidHODs/tsaw/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// NewHandler sets the repositories for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// performs a logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again.."

	// sends the data to the template
	render.RenderTemplate(w, "about_page.html", &models.TemplateData{
		StringMap: stringMap,
	}) 
}