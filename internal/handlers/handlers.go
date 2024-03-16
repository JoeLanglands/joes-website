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

func (repo *Repository) Root() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo.cfg.AddUniqueVisitor(r.RemoteAddr)
		repo.cfg.Logger.Info("served root page")
		repo.rdr.RenderTemplateWithComponents(w, r, "base.html", &models.TemplateData{
			IntMap: map[string]int{
				"unique_visitors": repo.cfg.GetUniqueVisitors(),
			},
		})
	})
}

func (repo *Repository) Title() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo.cfg.RequestColour <- struct{}{}

		colour := <-repo.cfg.TitleColourState

		repo.rdr.RenderTemplate(w, r, "title.html", &models.TemplateData{
			StringMap: map[string]string{
				"colour": colour,
			},
		})
	})
}

func (repo *Repository) Home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		repo.cfg.Logger.Info("served home page")
		repo.rdr.RenderTemplateWithComponents(w, r, "home.html", nil)
	})
}

func (repo *Repository) About() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		birthday := time.Date(1992, time.August, 11, 8, 0, 0, 0, time.FixedZone("GMT", 1))
		age := time.Since(birthday)
		repo.cfg.Logger.Info("served about page")
		repo.rdr.RenderTemplate(w, r, "about.html", &models.TemplateData{
			IntMap: map[string]int{
				"age": int(age.Seconds()),
			},
		})
	})
}

func (repo *Repository) Projects() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo.cfg.Logger.Info("served projects page")
		repo.rdr.RenderTemplate(w, r, "projects.html", nil)
	})
}

func (repo *Repository) Carousel() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo.cfg.RequestState <- struct{}{}

		// copying the sync.RWMutex here but we're not touching it after here so maybe its ok?
		carouselState := <-repo.cfg.CarouselState

		intMap := carouselState.Margin
		intMap["delay"] = state.CarouselPeriod

		repo.rdr.RenderTemplate(w, r, "carouselcontent.html", &models.TemplateData{
			StringMap: carouselState.Photo,
			IntMap:    carouselState.Margin,
		})
	})
}
