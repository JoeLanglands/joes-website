package main

import (
	"io/fs"
	"net/http"
	"net/http/pprof"

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

	mux.Get("/", handlers.Repo.Root())
	mux.Get("/home", router.Use(router.OnlyServeHTMX, handlers.Repo.Home()))
	mux.Get("/about", router.Use(router.OnlyServeHTMX, handlers.Repo.About()))
	mux.Get("/projects", router.Use(router.OnlyServeHTMX, handlers.Repo.Projects()))
	mux.Get("/carousel", router.Use(router.OnlyServeHTMX, handlers.Repo.Carousel()))
	mux.Get("/title", router.Use(router.OnlyServeHTMX, handlers.Repo.Title()))

	// add pprof routes
	mux.GetFunc("/debug/pprof/", pprof.Index)
	mux.GetFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.GetFunc("/debug/pprof/profile", pprof.Profile)
	mux.GetFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.GetFunc("/debug/pprof/trace", pprof.Trace)

	return mux
}
