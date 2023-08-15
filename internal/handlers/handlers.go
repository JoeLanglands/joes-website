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

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// TODO doesn't like this below, you need to figure out where log goes
	repo.cfg.Logger.Info("Home page accessed")
	repo.rdr.RenderTemplate(w, r, "base.gohtml", nil)
}
