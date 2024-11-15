package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

func init() {
	appsCmd.AddCommand(appsListCmd, appsOpenCmd)

	rootCmd.AddCommand(appsCmd)

	// Define flags for the apps open command
	appsOpenCmd.Flags().StringP("uri", "u", "", "URI of the app to open")
	appsOpenCmd.Flags().StringP("name", "n", "", "Name of the app to open")
}

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List and open apps on your TV",
	Long:  `Allows you to list and open apps on your TV.`,
}

var appsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List apps on your TV",
	Run: func(cmd *cobra.Command, args []string) {
		result, _, err := client.AppControl.GetApplicationList()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		apps := result.Result[0]

		for i, app := range apps {
			fmt.Printf("%2d. %s\n", i+1, app.Title)
		}
	},
}

var appsOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an app on your TV",
	Long:  `Allows you to open app using the URI or name of the app.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Validate only one of --uri or --name is provided
		if cmd.Flags().Changed("uri") == cmd.Flags().Changed("name") {
			return fmt.Errorf("eather --uri or --name must be provided")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var uri string

		if cmd.Flags().Changed("uri") {
			value, err := cmd.Flags().GetString("uri")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				os.Exit(1)
			}
			uri = value
		} else {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				os.Exit(1)
			}

			// Get list of apps
			result, _, err := client.AppControl.GetApplicationList()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				os.Exit(1)
			}
			apps := result.Result[0]

			// Gather all app titles to perform fuzzy search
			var titles []string
			appMap := make(map[string]string, len(apps))
			for _, app := range apps {
				titles = append(titles, app.Title)
				appMap[app.Title] = app.URI
			}

			// Perform fuzzy search for closest match
			matches := fuzzy.RankFindFold(name, titles)
			if len(matches) == 0 {
				fmt.Fprintf(os.Stderr, "No matching app found for name: %s\n", name)
				os.Exit(1)
			}
			sort.Sort(matches)

			// Use the first match for now
			matchedTitle := matches[0].Target
			uri = appMap[matchedTitle]
			fmt.Printf("Found app: %s (URI: %s)\n", matchedTitle, uri)
		}

		_, _, err := client.AppControl.SetActiveApp(uri, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	},
}
