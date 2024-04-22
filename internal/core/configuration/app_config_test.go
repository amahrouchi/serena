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
	app := tests.NewTestApp(false).Run(s.T(), fx.Populate(&config))
	defer app.RequireStop()

	s.Equal("test", config.App.Env)
	s.NotEmpty(config.App.Port)
	s.NotEmpty(config.App.Db.Host)
	s.NotEmpty(config.App.Db.Port)
	s.NotEmpty(config.App.Db.User)
	s.NotEmpty(config.App.Db.Password)
	s.NotEmpty(config.App.Db.DbName)
	s.False(config.App.BlockChain.WorkerEnabled)
	s.NotEmpty(config.App.BlockChain.Interval)
}

// TestNewConfig tests the NewConfig method.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}
