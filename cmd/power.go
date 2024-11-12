package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
		_, _, err := client.System.SetPowerStatus(true)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var powerOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off the TV",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, err := client.System.SetPowerStatus(false)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var powerStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the power status of the TV",
	Run: func(cmd *cobra.Command, args []string) {
		result, _, err := client.System.GetPowerStatus()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *result.Result != nil && len(*result.Result) > 0 {
			status := (*result.Result)[0].Status
			fmt.Println(status)
			return
		} else {
			fmt.Println("Failed to get power status")
			os.Exit(1)
		}
	},
}
