package main

import (
	"io/fs"
	"log/slog"
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

func getRouter(log *slog.Logger) http.Handler {
	mux := router.NewMux(router.WithLogger(log))
	fsys := getStaticFS()

	fileserver := NewFileServer(http.FileServer(fsys))

	mux.Handle("/static/", fileserver)

	logHtmxStack := router.UseStack(router.OnlyServeHTMX, router.RequestLogging)

	mux.Get("/", router.Use(router.RequestLogging, handlers.Repo.Root()))
	mux.Get("/home", logHtmxStack(handlers.Repo.Home()))
	mux.Get("/about", logHtmxStack(handlers.Repo.About()))
	mux.Get("/projects", logHtmxStack(handlers.Repo.Projects()))
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
