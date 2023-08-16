package handlers

import (
	"net/http"

	"github.com/JoeLanglands/joes-website/internal/config"
	"github.com/JoeLanglands/joes-website/internal/render"
)

type Repository struct {
	cfg *config.SiteConfig
	rdr render.Renderer
}

var Repo *Repository

func NewRepo(cfg *config.SiteConfig) *Repository {
	return &Repository{
		cfg: cfg,
		rdr: render.NewRenderer(cfg),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Root(w http.ResponseWriter, r *http.Request) {
	repo.rdr.RenderTemplate(w, r, "base.gohtml", nil)
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	repo.rdr.RenderComponent(w, r, "home.gohtml", nil)
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	repo.rdr.RenderComponent(w, r, "about.gohtml", nil)
}

func (repo *Repository) Projects(w http.ResponseWriter, r *http.Request) {
	repo.rdr.RenderComponent(w, r, "projects.gohtml", nil)
}
