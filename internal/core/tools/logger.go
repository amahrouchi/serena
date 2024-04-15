package tools

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"os"
)

// NewLogger creates a new logger.
func NewLogger(config *configuration.Config) *zerolog.Logger {
	level := lo.Ternary(config.Env == configuration.EnvDev, zerolog.DebugLevel, zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()

	return &logger
}
