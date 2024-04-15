package tests

import "github.com/rs/zerolog"

// NewEmptyLogger creates a new empty logger.
func NewEmptyLogger() *zerolog.Logger {
	logger := zerolog.New(nil).Level(zerolog.Disabled)

	return &logger
}
