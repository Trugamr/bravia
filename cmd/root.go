package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/trugamr/bravia-cli/bravia"
)

// client is the common client used for all commands
var client *bravia.Client

func init() {
	// Create common client for all commands
	baseURL, err := url.Parse(os.Getenv("BRAVIA_BASE_URL"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	psk := os.Getenv("BRAVIA_PSK")
	client = bravia.NewClient(baseURL).WithAuthPSK(psk)
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
