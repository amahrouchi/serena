package configuration

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

// Config represents the application configuration.
type Config struct {
	Env           string
	Port          int
	BlockDuration int
	DbDsn         string
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
	viper.SetDefault("env", EnvProd)
	viper.SetDefault("port", 8080)
	viper.SetDefault("blockDuration", 300)

	// Bind environment variables to Viper
	_ = viper.BindEnv("env", "SRN_ENV")
	_ = viper.BindEnv("port", "SRN_PORT")
	_ = viper.BindEnv("blockDuration", "SRN_BLOCK_DURATION")
	_ = viper.BindEnv("dbDsn", "SRN_DB_DSN")

	// Unmarshal the configuration
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}

// LoadConfig loads the configuration.
func LoadConfig(config *Config, logger *zerolog.Logger) error {
	logger.Info().Interface("config", config).Msgf("The config has been loaded")

	return nil
}
