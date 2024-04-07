package main

import (
	"github.com/JoeLanglands/joes-website/pkg/jmux"
	"github.com/JoeLanglands/joes-website/pkg/jmux/middlewares"
	"io/fs"
	"log/slog"
	"net/http"
	"net/http/pprof"

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

func getRouter(log *slog.Logger) http.Handler {
	mux := jmux.NewMux(jmux.WithLogger(log))
	fsys := getStaticFS()

	fileserver := NewFileServer(http.FileServer(fsys))

	mux.Handle("/static/", fileserver)

	onlyServeHtmxMw := middlewares.OnlyServeHTMXHandler(handlers.Repo.OnlyHTMXHandler())

	logHtmxStack := jmux.UseStack(onlyServeHtmxMw, middlewares.RequestLogging)

	mux.Get("/{$}", jmux.Use(middlewares.RequestLogging, handlers.Repo.Root()))
	mux.Get("/home", logHtmxStack(handlers.Repo.Home()))
	mux.Get("/about", logHtmxStack(handlers.Repo.About()))
	mux.Get("/projects", logHtmxStack(handlers.Repo.Projects()))
	mux.Get("/carousel", jmux.Use(onlyServeHtmxMw, handlers.Repo.Carousel()))
	mux.Get("/title", jmux.Use(onlyServeHtmxMw, handlers.Repo.Title()))

	// add pprof routes
	mux.GetFunc("/debug/pprof/", pprof.Index)
	mux.GetFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.GetFunc("/debug/pprof/profile", pprof.Profile)
	mux.GetFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.GetFunc("/debug/pprof/trace", pprof.Trace)

	mux.NotFoundHandler(handlers.Repo.NotFound())

	return mux
}
