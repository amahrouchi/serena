package core

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

// Config represents the application configuration.
type Config struct {
	logger *zerolog.Logger

	Env  string
	Port int
}

// NewConfig creates a new Config.
func NewConfig(logger *zerolog.Logger) *Config {
	config := Config{
		logger: logger,
	}
	config.init()

	return &config
}

// init initializes the configuration.
func (c *Config) init() {
	c.logger.Info().Msg("Initializing configuration from environment...")

	// Load the configuration file
	viper.AutomaticEnv()

	// Bind environment variables to Viper
	_ = viper.BindEnv("env", "SRN_ENV")
	_ = viper.BindEnv("port", "SRN_PORT")

	// Unmarshal the configuration
	err := viper.Unmarshal(&c)
	c.logger.Info().Msgf("Configuration: %+v", c)
	if err != nil {
		c.logger.Error().Err(err).Msg("Unable to load configuration from environment...")
		panic(err)
	}

	c.logger.Info().Msg("Configuration initialized")
}
