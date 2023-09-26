package main

import (
	"io/fs"
	"net/http"

	"github.com/JoeLanglands/joes-website/internal/handlers"
	"github.com/JoeLanglands/joes-website/internal/router"
	"github.com/JoeLanglands/joes-website/static"
)

func getStaticFS() http.FileSystem {
	f, err := fs.Sub(static.StaticFS, ".")
	if err != nil {
		panic(err)
	}
	return http.FS(f)
}

func getRouter() http.Handler {
	mux := router.NewMux()
	fs := getStaticFS()

	fileserver := NewFileServer(http.FileServer(fs))

	mux.Handle("/static/", fileserver)

	mux.Get("/", handlers.Repo.Root)
	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/projects", handlers.Repo.Projects)
	mux.Get("/carousel", handlers.Repo.Carousel)
	mux.Get("/title", handlers.Repo.Title)

	return mux
}
