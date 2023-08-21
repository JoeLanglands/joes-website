package config

import (
	"log/slog"

	"github.com/JoeLanglands/joes-website/internal/state"
)

const (
	configFile = "config.json"
)

type SiteConfig struct {
	InProduction   bool                     `json:"in_production"`
	CarouselPeriod int                      `json:"carousel_period"`
	Logger         *slog.Logger             `json:"-"`
	Msg            chan []byte              `json:"-"`
	CarouselState  chan state.CarouselState `json:"-"`
	RequestState   chan struct{}            `json:"-"`
}

func NewConfig() *SiteConfig {
	return &SiteConfig{}
}

func LoadOrNewConfig() *SiteConfig {
	return &SiteConfig{}
}

func (cfg *SiteConfig) SaveConfig() {

}
