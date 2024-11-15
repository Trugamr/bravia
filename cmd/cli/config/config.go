package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL string `mapstructure:"base_url"`
	PSK     string `mapstructure:"psk"`
}

func New() *Config {
	return &Config{}
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
	// Check: https://github.com/spf13/viper/issues/188#issuecomment-255519149
	viper.BindEnv("BASE_URL")
	viper.BindEnv("PSK")

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

	return nil
}

// AddFlags adds flags to a Cobra command and binds them to Viper
func (c *Config) AddFlags(cmd *cobra.Command) {
	// Define flags
	cmd.PersistentFlags().StringVar(&c.BaseURL, "base-url", "", "Base URL for the API")
	cmd.PersistentFlags().StringVar(&c.PSK, "psk", "", "Pre-shared key for the API")

	// Bind flags to Viper
	viper.BindPFlag("base_url", cmd.Flags().Lookup("base-url"))
	viper.BindPFlag("psk", cmd.Flags().Lookup("psk"))
}
