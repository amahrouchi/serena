package configuration

import (
	"errors"
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

// LoadConfig loads the configuration.
func LoadConfig(config *Config, logger *zerolog.Logger) error {
	logger.Debug().
		Interface("config", config).
		Msgf("The config has been loaded")

	return nil
}

// --------------------------------------------
// --------------------------------------------
// --------------------------------------------

// Config represents the whole application configuration.
type Config struct {
	App *AppConfig `mapstructure:"app"`
}

// AppConfig represents the application specific configuration.
type AppConfig struct {
	Env        string            `mapstructure:"env"`
	Port       int               `mapstructure:"port"`
	BlockChain *BlockChainConfig `mapstructure:"blockchain"`
	Db         *DbConfig         `mapstructure:"db"`
}

// BlockChainConfig represents the blockchain specific configuration.
type BlockChainConfig struct {
	WorkerEnabled bool `mapstructure:"worker_enabled"`
	Interval      int  `mapstructure:"interval"`
}

// DbConfig represents the database specific configuration.
type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
}

// NewConfig creates a new Config.
func NewConfig() *Config {
	// Find the configuration path
	configPath, err := getConfigPath()
	if err != nil {
		panic(err)
	}

	// Load the configuration
	config := &Config{}
	config.init(configPath, "config", false)

	// Overload with env files
	env := os.Getenv("SRN_ENV")
	envConfig := "config." + env
	stats, err := os.Stat(configPath + "/" + envConfig + ".yml")
	if err == nil && !stats.IsDir() {
		config.init(configPath, envConfig, true)
	}
	return config
}

// init initializes the configuration.
func (c *Config) init(configPath, configName string, reset bool) {
	// Reset the configuration
	if reset {
		viper.Reset()
	}

	// Load the configuration file
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
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
				viper.Set(key, envVar)
			}
		}
	}

	// Unmarshal the configuration into the struct
	err := viper.Unmarshal(c)
	if err != nil {
		panic(err)
	}
}

// getConfigPath retrieves the configuration path from the current directory or the parent directories.
func getConfigPath() (string, error) {
	const configDirName = "configs"

	// Check the current directory
	stats, err := os.Stat("./" + configDirName)
	if err == nil && stats.IsDir() {
		return "./" + configDirName, nil
	}

	// Check 10 parent directories above
	backPath := ""
	for i := 1; i <= 10; i++ {
		backPath += "../"
		stats, err := os.Stat(backPath + configDirName)
		if err == nil && stats.IsDir() {
			return backPath + configDirName, nil
		}
	}

	return "", errors.New("could not find the config directory")
}
