package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	BaseURL string `mapstructure:"base_url"`
	PSK     string `mapstructure:"psk"`
	Port    string `mapstructure:"port"`
}

func New() *Config {
	return &Config{
		Port: "8080", // Default port
	}
}

// Load initializes Viper, loads the configuration, and populates the Config struct
func (c *Config) Load() error {
	// Set up config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Add config paths
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Join(home, ".bravia"))

	// Set environment variable key overrides
	viper.SetEnvPrefix("BRAVIA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Manually bind environment variables
	viper.BindEnv("BASE_URL")
	viper.BindEnv("PSK")
	viper.BindEnv("PORT")

	// Attempt to read the config file, ignore error if not found
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error loading config file: %w", err)
		}
	}

	// Unmarshal into the Config struct
	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Validate required fields
	if c.BaseURL == "" {
		return fmt.Errorf("base_url is required (set via config file, --base-url flag, or BRAVIA_BASE_URL env var)")
	}
	if c.PSK == "" {
		return fmt.Errorf("psk is required (set via config file, --psk flag, or BRAVIA_PSK env var)")
	}

	return nil
}
