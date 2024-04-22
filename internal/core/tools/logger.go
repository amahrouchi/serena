package tools

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"os"
)

// NewLogger creates a new logger.
func NewLogger(config *configuration.Config) *zerolog.Logger {
	level := lo.Ternary(config.App.Env == configuration.EnvProd, zerolog.InfoLevel, zerolog.DebugLevel)
	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()

	return &logger
}
