package cmds

import (
	"fmt"

	"github.com/kuhahalong/ddgsearch"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Perform a DuckDuckGo text search",
	Example: `  ddgs search "golang programming"
  ddgs search -n 5 -r us-en "golang tutorial"
  ddgs search --output results.json "web development"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getClient()
		if err != nil {
			return fmt.Errorf("failed to create client: %w", err)
		}

		params := &ddgsearch.SearchParams{
			Query:      args[0],
			MaxResults: maxResults,
			Region:     ddgsearch.Region(region),
			SafeSearch: ddgsearch.SafeSearch(safeSearch),
		}

		results, err := client.Search(cmd.Context(), params)
		if err != nil {
			return fmt.Errorf("search failed: %w", err)
		}

		if outputFile != "" {
			if err := saveResults(results); err != nil {
				return fmt.Errorf("failed to save results: %w", err)
			}
		} else {
			printResults(results)
		}
		return nil
	},
}
