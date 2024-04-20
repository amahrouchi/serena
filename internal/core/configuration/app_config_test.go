package configuration_test

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"testing"
)

// ConfigSuite tests the Config.
type ConfigSuite struct {
	suite.Suite
}

// TestLoadConfig tests the LoadConfig method.
func (s *ConfigSuite) TestLoadConfig() {
	var config *configuration.Config
	tests.NewTestApp(false).Run(s.T(), fx.Populate(&config))

	s.Equal("test", config.Env)
	s.NotEmpty(config.Port)
	s.NotEmpty(config.DbHost)
	s.NotEmpty(config.DbPort)
	s.NotEmpty(config.DbUser)
	s.NotEmpty(config.DbPassword)
	s.NotEmpty(config.DbName)
	s.False(config.BlockWorkerEnabled)
	s.NotEmpty(config.BlockDuration)
}

// TestNewConfig tests the NewConfig method.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}
