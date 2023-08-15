package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/JoeLanglands/joes-website/internal/config"
	"github.com/JoeLanglands/joes-website/internal/handlers"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				newTime := a.Value.Time().Format("2006-01-02 15:04:05Z")
				a.Value = slog.StringValue(newTime)
				return slog.Attr{
					Key:   "time",
					Value: a.Value,
				}
			}
			return a
		},
	})

	logger := slog.New(jsonHandler)
	cfg := &config.SiteConfig{
		InProduction: false,
		Logger:       logger,
	}

	repo := handlers.NewRepo(cfg)
	handlers.NewHandlers(repo)

	mux := router()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
