package render

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/JoeLanglands/joes-website/internal/config"
)

//go:embed templates/*
var fs embed.FS

type Renderer struct {
	cfg *config.SiteConfig
}

func NewRenderer(cfg *config.SiteConfig) Renderer {
	return Renderer{cfg: cfg}
}

func (rdr *Renderer) RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data any) error {
	buf := new(bytes.Buffer)

	nameGlob := fmt.Sprintf("templates/%s", name)

	tmpl, err := template.ParseFS(fs, nameGlob, "templates/*.layout.gohtml")
	if err != nil {
		rdr.cfg.Logger.Error("error parsing template", "error", err)
		return err
	}

	err = tmpl.Execute(buf, data)
	if err != nil {
		rdr.cfg.Logger.Error("error executing template", "error", err)
		return err
	}

	n, err := buf.WriteTo(w)
	if err != nil {
		rdr.cfg.Logger.Error("error writing template to response writer", "error", err)
		return err
	}
	rdr.cfg.Logger.Info("rendered template", "name", name, "bytes", n)

	return nil
}

func (rdr *Renderer) RenderComponent(w http.ResponseWriter, r *http.Request, name string, data any) error {
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

	n, err := buf.WriteTo(w)
	if err != nil {
		rdr.cfg.Logger.Error("error writing template to response writer", "error", err)
		return err
	}
	rdr.cfg.Logger.Info("rendered template", "name", name, "bytes", n)

	return nil
}
