package core

import (
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"os"
)

// newLogger creates a new logger.
func newLogger(config *Config) *zerolog.Logger {
	level := lo.Ternary(config.Env == EnvDev, zerolog.DebugLevel, zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()

	return &logger
}

// NewEmptyLogger creates a new empty logger.
func NewEmptyLogger() *zerolog.Logger {
	logger := zerolog.New(nil).Level(zerolog.Disabled)

	return &logger
}
