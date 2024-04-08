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
	Env  string
	Port int
}

// NewConfig creates a new Config.
func NewConfig() *Config {
	config := &Config{}
	config.init()

	return config
}

// init initializes the configuration.
func (c *Config) init() {
	// Load the configuration file
	viper.AutomaticEnv()

	// Set the default values
	viper.SetDefault("env", envProd)
	viper.SetDefault("port", 8080)

	// Bind environment variables to Viper
	_ = viper.BindEnv("env", "SRN_ENV")
	_ = viper.BindEnv("port", "SRN_PORT")

	// Unmarshal the configuration
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}

func LoadConfig(config *Config, logger *zerolog.Logger) {
	logger.Info().Interface("config", config).Msgf("The config has been loaded")
}
