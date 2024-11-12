package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/trugamr/bravia-cli/bravia"
	"github.com/trugamr/bravia-cli/config"
)

var (
	client *bravia.Client
	cfg    *config.Config
)

func init() {
	cfg = config.New()

	// Add config flags to root command
	cfg.AddFlags(rootCmd)

	// Initialize configuration before any command runs
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if err := cfg.Load(); err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// Create client using config values
	baseURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		fmt.Println("Invalid base URL:", err)
		os.Exit(1)
	}

	client = bravia.NewClient(baseURL).WithAuthPSK(cfg.PSK)
}

var rootCmd = &cobra.Command{
	Use:   "bravia",
	Short: "Control your Sony Bravia TV from the command line",
	Long: `A CLI tool for managing your Sony Bravia TV. 
Allows you to control volume, switch inputs, launch apps, and perform other remote functions 
through simple commands.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
