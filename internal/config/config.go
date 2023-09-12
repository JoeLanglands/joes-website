package config

import (
	"encoding/json"
	"log/slog"
	"net"
	"os"

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

// NewConfig returns a (pointer to) new default SiteConfig
func NewConfig() *SiteConfig {
	return &SiteConfig{
		UniqueVisitors: make(map[string]struct{}),
	}
}

func LoadOrNewConfig() *SiteConfig {
	cfgData, err := os.ReadFile(configFile)
	if err != nil {
		slog.Error("Failed to read config file, using defaults", "error", err)
		return NewConfig()
	}
	var cfg SiteConfig
	err = json.Unmarshal(cfgData, &cfg)
	if err != nil {
		slog.Error("Failed to unmarshal config file, using defaults", "error", err)
		return NewConfig()
	}
	return &cfg
}

func (cfg *SiteConfig) GetUniqueVisitors() int {
	return len(cfg.UniqueVisitors)
}

func (cfg *SiteConfig) AddUniqueVisitor(remoteAddr string) {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		cfg.Logger.Error("Failed to split host and port", "error", err)
	}
	cfg.UniqueVisitors[host] = struct{}{}
}

func (cfg *SiteConfig) SaveConfig() {
	cfgData, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		cfg.Logger.Error("Failed to marshal config", "error", err)
		return
	}
	err = os.WriteFile(configFile, cfgData, 0644)
	if err != nil {
		cfg.Logger.Error("Failed to write config", "error", err)
	}
}
