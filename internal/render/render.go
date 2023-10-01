package render

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/JoeLanglands/joes-website/internal/config"
	"github.com/JoeLanglands/joes-website/internal/models"
)

//go:embed templates/*
var fs embed.FS

type Renderer struct {
	cfg *config.SiteConfig
}

func NewRenderer(cfg *config.SiteConfig) Renderer {
	return Renderer{cfg: cfg}
}

// RenderTemplateWithComponents renders the template given by name and along with all *.component.gohtml files
func (rdr *Renderer) RenderTemplateWithComponents(w http.ResponseWriter, r *http.Request, name string, data any) error {
	buf := new(bytes.Buffer)

	nameGlob := fmt.Sprintf("templates/%s", name)

	tmpl, err := template.ParseFS(fs, nameGlob, "templates/*.component.gohtml")
	if err != nil {
		rdr.cfg.Logger.Error("error parsing template", "error", err)
		return err
	}

	err = tmpl.Execute(buf, data)
	if err != nil {
		rdr.cfg.Logger.Error("error executing template", "error", err)
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		rdr.cfg.Logger.Error("error writing template to response writer", "error", err)
		return err
	}

	return nil
}

// RenderTemplate renders a single template as a component to the response writer mainly for htmx elements.
func (rdr *Renderer) RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data *models.TemplateData) error {
	buf := new(bytes.Buffer)

	nameGlob := fmt.Sprintf("templates/%s", name)

	tmpl, err := template.ParseFS(fs, nameGlob)
	if err != nil {
		rdr.cfg.Logger.Error("error parsing template", "error", err)
		return err
	}

	err = tmpl.Execute(buf, data)
	if err != nil {
		rdr.cfg.Logger.Error("error executing template", "error", err)
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		rdr.cfg.Logger.Error("error writing template to response writer", "error", err)
		return err
	}

	return nil
}
