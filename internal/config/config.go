package config

import (
	"encoding/json"
	"log/slog"
	"os"
	"strings"

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
	UniqueVisitors map[string]struct{}      `json:"unique_visitors"`
}

func LoadOrNewConfig() *SiteConfig {
	return &SiteConfig{
		UniqueVisitors: make(map[string]struct{}),
	}
}

func (cfg *SiteConfig) GetUniqueVisitors() int {
	return len(cfg.UniqueVisitors)
}

func (cfg *SiteConfig) AddUniqueVisitor(fullIP string) {
	ip, _, ok := strings.Cut(fullIP, ":")
	if !ok {
		cfg.Logger.Info("Failed to cut IP address", "ip", fullIP)
	}
	cfg.UniqueVisitors[ip] = struct{}{}
}

func (cfg *SiteConfig) SaveConfig() {
	cfgData, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		cfg.Logger.Error("Failed to marshal config", "error", err)
		return
	}
	err = os.WriteFile("site_config.json", cfgData, 0644)
	if err != nil {
		cfg.Logger.Error("Failed to write config", "error", err)
	}
}
