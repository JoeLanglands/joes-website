package config

import (
	"log/slog"

	"github.com/JoeLanglands/joes-website/internal/state"
)

type SiteConfig struct {
	InProduction  bool
	Logger        *slog.Logger
	Msg           chan []byte
	CarouselState chan state.CarouselState
	RequestState  chan struct{}
}
