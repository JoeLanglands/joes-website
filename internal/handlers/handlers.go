package handlers

import (
	"net/http"
	"time"

	"github.com/JoeLanglands/joes-website/internal/config"
	"github.com/JoeLanglands/joes-website/internal/models"
	"github.com/JoeLanglands/joes-website/internal/render"
	"github.com/JoeLanglands/joes-website/internal/state"
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
	repo.cfg.AddUniqueVisitor(r.RemoteAddr)
	repo.rdr.RenderTemplateWithComponents(w, r, "base.gohtml", &models.TemplateData{
		IntMap: map[string]int{
			"unique_visitors": repo.cfg.GetUniqueVisitors(),
		},
	})
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	repo.rdr.RenderTemplateWithComponents(w, r, "home.gohtml", nil)
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	birthday := time.Date(1992, time.August, 11, 8, 0, 0, 0, time.FixedZone("GMT", 1))
	age := time.Since(birthday)
	repo.rdr.RenderTemplate(w, r, "about.gohtml", &models.TemplateData{
		IntMap: map[string]int{
			"age": int(age.Seconds()),
		},
	})
}

func (repo *Repository) Projects(w http.ResponseWriter, r *http.Request) {
	repo.rdr.RenderTemplate(w, r, "projects.gohtml", nil)
}

func (repo *Repository) Carousel(w http.ResponseWriter, r *http.Request) {
	repo.cfg.RequestState <- struct{}{}

	// copying the sync.RWMutex here but we're not touching it after here so maybe its ok?
	carouselState := <-repo.cfg.CarouselState

	intMap := carouselState.Margin
	intMap["delay"] = state.CarouselPeriod

	repo.rdr.RenderTemplate(w, r, "carouselcontent.gohtml", &models.TemplateData{
		StringMap: carouselState.Photo,
		IntMap:    carouselState.Margin,
	})
}
