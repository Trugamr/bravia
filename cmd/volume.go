package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// Add volume command to root command
	rootCmd.AddCommand(volumeCmd)

	// Define flags for the volume command
	volumeCmd.Flags().StringP("level", "l", "", "Set the volume level (e.g., +10, -5, 24)")
	volumeCmd.Flags().StringP("target", "t", "", "Set the target (e.g., speaker, headphone)")

	// Mark required flags
	volumeCmd.MarkFlagRequired("level")
}

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Control the volume of the TV",
	Long:  `Allows setting the volume of the TV.`,
	Run: func(cmd *cobra.Command, args []string) {
		level, err := cmd.Flags().GetString("level")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		target, err := cmd.Flags().GetString("target")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		_, _, err = client.Audio.SetAudioVolume(level, target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	},
}
