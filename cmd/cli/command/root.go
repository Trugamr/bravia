package command

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/trugamr/bravia/api"
	"github.com/trugamr/bravia/cmd/cli/config"
)

var (
	client *api.Client
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
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	// Create client using config values
	baseURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	client = api.NewClient(baseURL).WithAuthPSK(cfg.PSK)
}

var rootCmd = &cobra.Command{
	Use:   "bravia",
	Short: "Control your Sony Bravia TV from the command line",
	Long: `A CLI tool for managing your Sony Bravia TV. 
Allows you to control volume, switch inputs, launch apps, and perform other remote functions 
through simple commands.`,
}

// ExecuteRoot is the entrypoint for the CLI
func ExecuteRoot() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
