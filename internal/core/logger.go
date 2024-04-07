package core

import (
	"github.com/rs/zerolog"
	"os"
)

func NewLogger() *zerolog.Logger {
	// TODO: Make the level vary based on the environment.
	level := zerolog.DebugLevel
	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()

	return &logger
}
