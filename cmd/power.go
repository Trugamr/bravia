package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/trugamr/bravia-cli/bravia"
)

func init() {
	powerCmd.AddCommand(powerOnCmd)
	powerCmd.AddCommand(powerOffCmd)
	powerCmd.AddCommand(powerStatusCmd)

	rootCmd.AddCommand(powerCmd)
}

var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "Control the power state of the TV",
	Long:  "Allows turning the TV on, off, or checking the power status.",
}

var powerOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn on the TV",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement command
		cmd.Help()
	},
}

var powerOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off the TV",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement command
		cmd.Help()
	},
}

var powerStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the power status of the TV",
	Run: func(cmd *cobra.Command, args []string) {
		baseURL, err := url.Parse(os.Getenv("BRAVIA_BASE_URL"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := bravia.NewClient(baseURL)

		status, err := client.System.GetPowerStatus()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(status)
	},
}
