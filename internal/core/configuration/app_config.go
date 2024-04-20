package configuration

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
)

const (
	EnvDev  = "dev"
	EnvTest = "test"
	EnvProd = "prod"
)

// Config represents the application configuration.
type Config struct {
	Env                string `mapstructure:"SRN_ENV"`
	Port               int    `mapstructure:"SRN_PORT"`
	BlockWorkerEnabled bool   `mapstructure:"SRN_BLOCK_WORKER_ENABLED"`
	BlockDuration      int    `mapstructure:"SRN_BLOCK_DURATION"`
	DbUser             string `mapstructure:"SRN_DB_USER"`
	DbPassword         string `mapstructure:"SRN_DB_PASSWORD"`
	DbHost             string `mapstructure:"SRN_DB_HOST"`
	DbPort             string `mapstructure:"SRN_DB_PORT"`
	DbName             string `mapstructure:"SRN_DB_NAME"`
}

// NewConfig creates a new Config.
func NewConfig() *Config {
	config := &Config{}
	config.init()

	return config
}

// init initializes the configuration.
func (c *Config) init() {
	// Load the environment variables
	viper.AutomaticEnv()

	// Load test environment variables
	configFile := ".env." + os.Getenv("SRN_ENV")
	viper.AddConfigPath("/app")
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Unmarshal the configuration
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}

// LoadConfig loads the configuration.
func LoadConfig(config *Config, logger *zerolog.Logger) error {
	logger.Debug().
		Interface("config", config).
		Msgf("The config has been loaded")

	return nil
}
