package config

import "log/slog"

type SiteConfig struct {
	InProduction bool
	Logger       *slog.Logger
}
