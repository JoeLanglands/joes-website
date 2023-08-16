package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/JoeLanglands/joes-website/internal/config"
	"github.com/JoeLanglands/joes-website/internal/handlers"
)

var cfg config.SiteConfig

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

	msgChan := make(chan []byte, 5)

	cfg.Logger = slog.New(jsonHandler)
	cfg.Msg = msgChan

	listenForMessages(&cfg)

	repo := handlers.NewRepo(&cfg)
	handlers.NewHandlers(repo)

	mux := getRouter()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
