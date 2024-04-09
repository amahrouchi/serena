package core

import (
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"os"
)

// NewLogger creates a new logger.
func NewLogger(config *Config) *zerolog.Logger {
	level := lo.Ternary(config.Env == EnvDev, zerolog.DebugLevel, zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()

	return &logger
}
