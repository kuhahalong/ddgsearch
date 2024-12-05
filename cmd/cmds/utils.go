package cmds

import (
	"github.com/kuhahalong/ddgsearch"
)

// getClient creates and returns a configured DuckDuckGo client
func getClient() (*ddgsearch.DDGS, error) {
	cfg := &ddgsearch.Config{
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		},
		Proxy:      proxyURL,
		Timeout:    timeout,
		MaxRetries: maxRetries,
		Cache:      true,
	}

	return ddgsearch.New(cfg)
}
