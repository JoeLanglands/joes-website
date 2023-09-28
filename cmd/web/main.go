package main

import (
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JoeLanglands/joes-website/internal/config"
	"github.com/JoeLanglands/joes-website/internal/handlers"
	"github.com/JoeLanglands/joes-website/internal/state"
)

var cfg = config.LoadOrNewConfig()

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

	sigExit := make(chan os.Signal)

	signal.Notify(sigExit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigExit
		cfg.SaveConfig()
		os.Exit(0)
	}()

	carouselChan := make(chan state.CarouselState)
	msgChan := make(chan []byte, 5)
	reqState := make(chan struct{})
	reqColour := make(chan struct{})
	titleColourChan := make(chan string)
	defer close(msgChan)
	defer close(carouselChan)
	defer close(reqState)
	defer close(titleColourChan)
	defer close(reqColour)

	cfg.Logger = slog.New(jsonHandler)
	cfg.Msg = msgChan
	cfg.CarouselState = carouselChan
	cfg.RequestState = reqState
	cfg.RequestColour = reqColour
	cfg.TitleColourState = titleColourChan

	state.Carouselhandler(carouselChan, reqState)
	state.TitleColourHandler(titleColourChan, reqColour)
	listenForMessages(cfg)

	repo := handlers.NewRepo(cfg)
	handlers.NewHandlers(repo)

	mux := getRouter()

	err := http.ListenAndServe(":6969", mux)
	if err != nil {
		panic(err)
	}
}
