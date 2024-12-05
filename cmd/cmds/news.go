package cmds

import (
	"fmt"

	"github.com/kuhahalong/ddgsearch"
	"github.com/spf13/cobra"
)

var newsCmd = &cobra.Command{
	Use:   "news [query]",
	Short: "Perform a DuckDuckGo news search",
	Example: `  ddgs news "latest technology"
  ddgs news --time-range day "crypto news"
  ddgs news -o news.csv "startup funding"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getClient()
		if err != nil {
			return fmt.Errorf("failed to create client: %w", err)
		}

		params := &ddgsearch.NewsParams{
			Query:      args[0],
			MaxResults: maxResults,
			Region:     ddgsearch.Region(region),
			SafeSearch: ddgsearch.SafeSearch(safeSearch),
			TimeRange:  ddgsearch.TimeRange(timeRange),
		}

		results, err := client.News(cmd.Context(), params)
		if err != nil {
			return fmt.Errorf("news search failed: %w", err)
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
