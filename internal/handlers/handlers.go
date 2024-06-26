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
		err := repo.rdr.RenderTemplateWithComponents(w, r, "base.html", &models.TemplateData{
			IntMap: map[string]int{
				"unique_visitors": repo.cfg.GetUniqueVisitors(),
			},
		})
		if err != nil {
			http.Error(w, "unable to render template :(", http.StatusInternalServerError)
		}
	})
}

func (repo *Repository) Title() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo.cfg.RequestColour <- struct{}{}

		colour := <-repo.cfg.TitleColourState

		err := repo.rdr.RenderTemplate(w, r, "title.html", &models.TemplateData{
			StringMap: map[string]string{
				"colour": colour,
			},
		})
		if err != nil {
			http.Error(w, "unable to render template :(", http.StatusInternalServerError)
		}
	})
}

func (repo *Repository) Home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := repo.rdr.RenderTemplateWithComponents(w, r, "home.html", nil)
		if err != nil {
			http.Error(w, "unable to render template :(", http.StatusInternalServerError)
		}
	})
}

func (repo *Repository) About() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		birthday := time.Date(1992, time.August, 11, 8, 0, 0, 0, time.FixedZone("GMT", 1))
		age := time.Since(birthday)

		err := repo.rdr.RenderTemplate(w, r, "about.html", &models.TemplateData{
			IntMap: map[string]int{
				"age": int(age.Seconds()),
			},
		})
		if err != nil {
			http.Error(w, "unable to render template :(", http.StatusInternalServerError)
		}
	})
}

func (repo *Repository) Projects() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := repo.rdr.RenderTemplate(w, r, "projects.html", nil)
		if err != nil {
			http.Error(w, "unable to render template :(", http.StatusInternalServerError)
		}
	})
}

func (repo *Repository) Carousel() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo.cfg.RequestState <- struct{}{}

		// copying the sync.RWMutex here, but we're not touching it after here so maybe it's ok?
		carouselState := <-repo.cfg.CarouselState

		intMap := carouselState.Margin
		intMap["delay"] = state.CarouselPeriod

		err := repo.rdr.RenderTemplate(w, r, "carouselcontent.html", &models.TemplateData{
			StringMap: carouselState.Photo,
			IntMap:    carouselState.Margin,
		})
		if err != nil {
			http.Error(w, "unable to render template :(", http.StatusInternalServerError)
		}
	})
}

func (repo *Repository) NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := repo.rdr.RenderTemplate(w, r, "notfound.html", nil)
		if err != nil {
			http.Error(w, "unable to render template :(", http.StatusInternalServerError)
		}
	})
}

func (repo *Repository) OnlyHTMXHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := repo.rdr.RenderTemplate(w, r, "htmxonly.html", nil)
		if err != nil {
			http.Error(w, "unable to render template :(", http.StatusInternalServerError)
		}
	})
}
