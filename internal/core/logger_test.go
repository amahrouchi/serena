package core

import (
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"testing"
)

// LoggerSuite tests the Logger.
type LoggerSuite struct {
	suite.Suite
}

// TestNewLogger tests the newLogger method.
func (s *LoggerSuite) TestNewLogger() {
	// Test the logger creation for dev environment
	s.Run("development", func() {
		config := &Config{Env: EnvDev}
		logger := newLogger(config)

		s.NotNil(logger)
		s.Equal(zerolog.DebugLevel, logger.GetLevel())
	})

	// Test the logger creation for prod environment
	s.Run("production", func() {
		config := &Config{Env: EnvProd}
		logger := newLogger(config)

		s.NotNil(logger)
		s.Equal(zerolog.InfoLevel, logger.GetLevel())
	})
}

// TestLoggerSuite tests the LoggerSuite.
func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}
