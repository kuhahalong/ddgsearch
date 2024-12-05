# DDGSearch

A native Go library for DuckDuckGo search functionality. This library provides a simple and efficient way to perform searches using DuckDuckGo's search engine.

## Why DuckDuckGo?

DuckDuckGo offers several advantages:
- **No Authentication Required**: Unlike other search engines, DuckDuckGo's API can be used without any API keys or authentication
- Privacy-focused search results
- No rate limiting for reasonable usage
- Support for multiple regions and languages
- Clean and relevant search results

## Features

- Clean and idiomatic Go implementation
- Comprehensive error handling
- Configurable search parameters
- In-memory caching with TTL
- Support for:
  - Multiple regions (us-en, uk-en, de-de, etc.)
  - Safe search levels (strict, moderate, off)
  - Time-based filtering (day, week, month, year)
  - Result pagination
  - Custom HTTP headers
  - Proxy configuration

## Installation

```bash
go get github.com/kuhahalong/ddgsearch
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/kuhahalong/ddgsearch"
)

func main() {
    // Create a new client with configuration
    cfg := &ddgsearch.Config{
        Timeout:    30 * time.Second,
        MaxRetries: 3,
        Cache:      true,
    }
    client, err := ddgsearch.New(cfg)
    if err != nil {
        log.Fatal(err)
    }

    // Configure search parameters
    params := &ddgsearch.SearchParams{
        Query:      "what is golang",
        Region:     ddgsearch.RegionUSEN,
        SafeSearch: ddgsearch.SafeSearchModerate,
        TimeRange:  ddgsearch.TimeRangeMonth,
        MaxResults: 10,
    }

    // Perform search
    response, err := client.Search(context.Background(), params)
    if err != nil {
        log.Fatal(err)
    }

    // Print results
    for i, result := range response.Results {
        fmt.Printf("%d. %s\n   URL: %s\n   Description: %s\n\n", 
            i+1, result.Title, result.URL, result.Description)
    }
}
```

## Advanced Usage

### Configuration

```go
// Create client with custom configuration
cfg := &ddgsearch.Config{
    Timeout:    20 * time.Second,
    MaxRetries: 3,
    Proxy:      "http://proxy:8080",
    Cache:      true,
    Headers: map[string]string{
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
    },
}
client, err := ddgsearch.New(cfg)
```

### Search Parameters

```go
params := &ddgsearch.SearchParams{
    Query:      "golang tutorial",       // Search query
    Region:     ddgsearch.RegionUSEN,    // Region for results (us-en, uk-en, etc.)
    SafeSearch: ddgsearch.SafeSearchModerate, // Safe search level
    TimeRange:  ddgsearch.TimeRangeWeek, // Time filter
    MaxResults: 10,                      // Maximum results to return
}
```

Available regions:
- RegionUSEN (United States)
- RegionUKEN (United Kingdom)
- RegionDEDE (Germany)
- RegionFRFR (France)
- RegionJPJP (Japan)
- RegionCNZH (China)
- RegionRURU (Russia)

Safe search levels:
- SafeSearchStrict
- SafeSearchModerate
- SafeSearchOff

Time range options:
- TimeRangeDay
- TimeRangeWeek
- TimeRangeMonth
- TimeRangeYear

### Proxy Support

```go
// HTTP proxy
cfg := &ddgsearch.Config{
    Proxy: "http://proxy:8080",
}
client, err := ddgsearch.New(cfg)

// SOCKS5 proxy
cfg := &ddgsearch.Config{
    Proxy: "socks5://proxy:1080",
}
client, err := ddgsearch.New(cfg)
```

## Command Line Interface

The library includes a command-line interface (CLI) tool for easy access to DuckDuckGo search functionality.

### Installation

```bash
go install github.com/kuhahalong/ddgsearch/cmd/ddgs@latest
```

### Usage

```bash
# Basic text search
ddgs search "golang programming"

# Search news
ddgs news "golang 1.21"
```

For more information, see [README.md](cmd/README.md)

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

### Development Setup

1. Clone the repository
```bash
git clone https://github.com/kuhahalong/ddgsearch.git
```

2. Install dependencies
```bash
go mod tidy
```

3. Run tests
```bash
go test ./...
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
