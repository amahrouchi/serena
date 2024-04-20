package tools_test

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"testing"
)

// LoggerSuite tests the Logger.
type LoggerSuite struct {
	suite.Suite
}

// TestNewLogger tests the NewLogger method.
func (s *LoggerSuite) TestNewLogger() {
	// Test the logger creation for dev environment
	s.Run("development", func() {
		config := &configuration.Config{Env: configuration.EnvDev}
		logger := tools.NewLogger(config)

		s.NotNil(logger)
		s.Equal(zerolog.DebugLevel, logger.GetLevel())
	})

	// Test the logger creation for prod environment
	s.Run("production", func() {
		config := &configuration.Config{Env: configuration.EnvProd}
		logger := tools.NewLogger(config)

		s.NotNil(logger)
		s.Equal(zerolog.InfoLevel, logger.GetLevel())
	})
}

// TestLoggerSuite tests the LoggerSuite.
func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}
