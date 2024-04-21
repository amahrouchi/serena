package configuration

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
	"regexp"
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
func LoadConfig(config *Config, configYaml *ConfigYaml, logger *zerolog.Logger) error {
	logger.Debug().
		Interface("config", config).
		Interface("configYaml", configYaml).
		Msgf("The config has been loaded")

	return nil
}

// --------------------------------------------
// --------------------------------------------
// --------------------------------------------

type ConfigYaml struct {
	App AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Env        string           `mapstructure:"env"`
	Port       int              `mapstructure:"port"`
	BlockChain BlockChainConfig `mapstructure:"blockchain"`
	Db         DbConfig         `mapstructure:"db"`
}

type BlockChainConfig struct {
	WorkerEnabled bool `mapstructure:"worker_enabled"`
	Interval      int  `mapstructure:"interval"`
}

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func NewConfigYaml() *ConfigYaml {
	config := &ConfigYaml{}
	config.init()

	return config
}

func (c *ConfigYaml) init() {
	// Load the configuration file
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/app/configs")
	// TODO: overload with env files
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Parse environment variables
	allKeys := viper.AllKeys()
	for _, key := range allKeys {
		value := viper.GetString(key)
		envRegex := `^\$\{([A-Z_]+)\}$`
		match := regexp.MustCompile(envRegex).FindStringSubmatch(value)
		if match != nil {
			envVar := os.Getenv(match[1])
			if envVar != "" {
				//if envVar == "true" || envVar == "false" {
				//	viper.Set(key, envVar == "true")
				//} else {
				viper.Set(key, envVar)
				//}
			}
		}
	}

	// Unmarshal the configuration into the struct
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}
