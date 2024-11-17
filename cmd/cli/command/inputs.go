package command

import (
	"cmp"
	"fmt"
	"os"
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
	"github.com/trugamr/bravia/api"
)

func init() {
	inputsCmd.AddCommand(inputsListCmd, inputsSelectCmd)

	rootCmd.AddCommand(inputsCmd)

	// Define flags for the inputs select command
	inputsSelectCmd.Flags().StringP("uri", "u", "", "URI of the input to select")
	inputsSelectCmd.Flags().StringP("name", "n", "", "Name of the input to select")
	inputsSelectCmd.Flags().StringP("label", "l", "", "Label of the input to select")
}

var inputsCmd = &cobra.Command{
	Use:   "inputs",
	Short: "List and control external inputs on your TV",
	Long:  `Allows you to list and control external inputs on your TV.`,
}

var inputsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List external inputs on your TV",
	Run: func(cmd *cobra.Command, args []string) {
		result, _, err := client.AVContent.GetCurrentExternalInputsStatus()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		inputs := result.Result[0]

		for i, input := range inputs {
			label := cmp.Or(input.Label, "-")
			fmt.Printf("%2d. %s (Label: %s) [URI: %s]\n", i+1, input.Title, label, input.URI)
		}
	},
}

var inputsSelectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select an external input on your TV",
	Long:  `Allows you to select an external input on your TV.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Validate only one of --uri, --name or --label is provided
		changes := [3]bool{cmd.Flags().Changed("uri"), cmd.Flags().Changed("name"), cmd.Flags().Changed("label")}
		changed := 0

		for _, change := range changes {
			if change {
				changed += 1
				break
			}
		}

		if changed != 1 {
			return fmt.Errorf("either --uri, --name or --label must be provided")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Helper to retrieve a flag value
		getFlagValue := func(flagName string) string {
			value, err := cmd.Flags().GetString(flagName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error retrieving flag %s: %s\n", flagName, err)
				os.Exit(1)
			}
			return value
		}

		// Helper to find the URI based on fuzzy matching
		findURI := func(input string, keySelector func(input api.ExternalInputStatus) string) string {
			result, _, err := client.AVContent.GetCurrentExternalInputsStatus()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching inputs: %s\n", err)
				os.Exit(1)
			}
			inputs := result.Result[0]

			// Gather keys and map them to URIs
			var keys []string
			inputMap := make(map[string]string, len(inputs))
			for _, input := range inputs {
				key := keySelector(input)
				keys = append(keys, key)
				inputMap[key] = input.URI
			}

			// Perform fuzzy search for closest match
			matches := fuzzy.RankFindFold(input, keys)
			if len(matches) == 0 {
				fmt.Fprintf(os.Stderr, "No matching input found for: %s\n", input)
				os.Exit(1)
			}
			sort.Sort(matches)

			matchedKey := matches[0].Target
			fmt.Printf("Found input: %s (URI: %s)\n", matchedKey, inputMap[matchedKey])
			return inputMap[matchedKey]
		}

		var uri string

		if cmd.Flags().Changed("uri") {
			uri = getFlagValue("uri")
		} else if cmd.Flags().Changed("name") {
			name := getFlagValue("name")
			uri = findURI(name, func(input api.ExternalInputStatus) string { return input.Title })
		} else {
			label := getFlagValue("label")
			uri = findURI(label, func(input api.ExternalInputStatus) string { return input.Label })
		}

		_, _, err := client.AVContent.SetPlayContent(uri)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Selected input:", uri)
	},
}
