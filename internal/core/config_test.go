package core

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// ConfigSuite tests the Config.
type ConfigSuite struct {
	suite.Suite
}

// TestLoadConfig tests the loadConfig method.
func (s *ConfigSuite) TestLoadConfig() {
	logger := NewEmptyLogger()
	err := loadConfig(&Config{}, logger)

	s.NoError(err)
}

// TestNewConfig tests the newConfig method.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}
