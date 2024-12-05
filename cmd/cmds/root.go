package cmds

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	// Global flags
	proxyURL    string
	timeout     time.Duration
	maxRetries  int
	outputFile  string
	outputType  string
	maxResults  int
	region      string
	safeSearch  string
	timeRange   string
	interactive bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "ddgs",
	Short:   "DuckDuckGo Search CLI",
	Version: "1.0.0",
}

func init() {
	// Global flags
	RootCmd.PersistentFlags().StringVarP(&proxyURL, "proxy", "p", "", "Proxy URL (e.g., http://proxy:8080)")
	RootCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 30*time.Second, "Request timeout")
	RootCmd.PersistentFlags().IntVarP(&maxRetries, "max-retries", "r", 3, "Maximum retry attempts")
	RootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Output file (json or csv)")
	RootCmd.PersistentFlags().StringVarP(&outputType, "type", "", "text", "Output type (text, json, csv)")
	RootCmd.PersistentFlags().IntVarP(&maxResults, "max-results", "n", 0, "Maximum number of results")
	RootCmd.PersistentFlags().StringVarP(&region, "region", "", "us-en", "Search region (us-en, uk-en, de-de, fr-fr, jp-jp, cn-zh, ru-ru)")
	RootCmd.PersistentFlags().StringVarP(&safeSearch, "safe-search", "", "moderate", "Safe search level (strict, moderate, off)")
	RootCmd.PersistentFlags().StringVarP(&timeRange, "time-range", "", "", "Time range (day, week, month, year)")
	RootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode")

	// Add commands
	RootCmd.AddCommand(searchCmd, newsCmd)
}
