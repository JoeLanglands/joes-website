package main

import (
	"io/fs"
	"net/http"

	"github.com/JoeLanglands/joes-website/internal/handlers"
	"github.com/JoeLanglands/joes-website/static"
)

func getStaticFS() http.FileSystem {
	f, err := fs.Sub(static.StaticFS, ".")
	if err != nil {
		panic(err)
	}
	return http.FS(f)
}

func router() *http.ServeMux {

	mux := http.NewServeMux()
	fs := getStaticFS()

	fileserver := NewFileServer(http.FileServer(fs))

	mux.HandleFunc("/", handlers.Repo.Home)
	mux.Handle("/static/", fileserver)

	return mux
}
