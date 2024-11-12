package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

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
