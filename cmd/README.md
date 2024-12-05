# DuckDuckGo Search CLI

A command-line interface for performing text and news searches using DuckDuckGo.

## Features

- Text search with customizable parameters
- News search with time range filtering
- Multiple output formats (text, JSON, CSV)
- Configurable search settings (region, safe search, etc.)
- Proxy support
- Interactive mode

## Installation

```bash
go install github.com/kuhahalong/ddgsearch/cmd/ddgs@latest
```

Or build from source:
```bash
git clone https://github.com/kuhahalong/ddgsearch
cd ddgsearch
go build -o ddgs ./cmd
```

## Usage

### Basic Commands

```bash
# Basic text search
ddgs search "golang programming"

# News search
ddgs search "latest tech news"

# Get help
ddgs --help
ddgs search --help
ddgs news --help
```

### Search Options

```bash
# Limit number of results
ddgs search -n 3 "golang tutorial"

# Change region
ddgs search --region uk-en "british news"

# Change safe search level
ddgs search --safe-search strict "images"

# Output as JSON
ddgs search --type json "programming"

# Save results to file
ddgs search -o results.json "web development"
ddgs search -o results.csv "machine learning"
```

### News Search Options

```bash
# News from last day
ddgs news --time-range day "cryptocurrency"

# News from last week with region
ddgs news --time-range week --region us-en "tech startups"

# Save news to CSV
ddgs news -o news.csv --type csv "startup funding"
```

### Advanced Options

```bash
# Use proxy
ddgs search --proxy http://proxy:8080 "search term"

# Change request timeout
ddgs search --timeout 60s "large query"

# Set maximum retries
ddgs search --max-retries 5 "unstable query"

# Interactive mode
ddgs search -i "step by step results"
```

## Available Flags

Global flags that apply to all commands:

- `-p, --proxy`: Proxy URL (e.g., http://proxy:8080)
- `-t, --timeout`: Request timeout (default: 30s)
- `-r, --max-retries`: Maximum retry attempts (default: 3)
- `-o, --output`: Output file (json or csv)
- `--type`: Output type (text, json, csv)
- `-n, --max-results`: Maximum number of results
- `--region`: Search region (us-en, uk-en, de-de, fr-fr, jp-jp, cn-zh, ru-ru)
- `--safe-search`: Safe search level (strict, moderate, off)
- `--time-range`: Time range for news (day, week, month, year)
- `-i, --interactive`: Interactive mode

## Testing Commands

Here are some test commands to verify different functionalities:

```bash
# Test basic functionality
ddgs search "test query"
ddgs news "test news"

# Test output formats
ddgs search --type json "test json"
ddgs search --type csv "test csv"
ddgs search -o test.json "test file output"
ddgs search -o test.csv "test csv output"

# Test search parameters
ddgs search --region us-en --safe-search strict "test params"
ddgs search -n 5 "test limit"
ddgs news --time-range day "test time range"

# Test error cases
ddgs search ""  # Empty query
ddgs search --region invalid "test invalid region"
ddgs search --safe-search invalid "test invalid safe search"
ddgs news --time-range invalid "test invalid time range"
```

## Error Handling

The CLI will provide clear error messages for common issues:

- Invalid parameters
- Network errors
- API errors
- File output errors

## Development

To modify or extend the CLI:

1. Main components:
   - `main.go`: Entry point and flag definitions
   - `search.go`: Text search implementation
   - `news.go`: News search implementation
   - `output.go`: Result formatting and file output
   - `utils.go`: Client configuration and utilities

2. Build from source:
```bash
cd /Users/bytedance/go/src/long/ddgsearch
go build -o ddgs ./cmd/ddgs
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
