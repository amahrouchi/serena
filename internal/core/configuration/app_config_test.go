package configuration

import (
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ConfigSuite tests the Config.
type ConfigSuite struct {
	suite.Suite
}

// TestLoadConfig tests the LoadConfig method.
func (s *ConfigSuite) TestLoadConfig() {
	logger := tests.NewEmptyLogger()
	err := LoadConfig(&Config{}, logger)

	s.NoError(err)
}

// TestNewConfig tests the NewConfig method.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}
