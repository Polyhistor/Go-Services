package handlers

import (
	"github.com/polyhistor/go-service/pkg/config"
	"github.com/polyhistor/go-service/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}


func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// NewHandlers sets the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

