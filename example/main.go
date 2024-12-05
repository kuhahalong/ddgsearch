package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kuhahalong/ddgsearch"
)

func main() {
	// Create a new client with custom configuration
	client, err := ddgsearch.New(&ddgsearch.Config{
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		},
		Timeout: 10 * time.Second,
		Cache:   true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create search parameters
	params := &ddgsearch.SearchParams{
		Query:      "golang programming",
		Region:     ddgsearch.RegionUS,
		SafeSearch: ddgsearch.SafeSearchModerate,
		TimeRange:  ddgsearch.TimeRangeMonth,
		MaxResults: 10,
	}

	// Perform search
	ctx := context.Background()
	response, err := client.Search(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	// Print results
	fmt.Printf("Found %d results:\n\n", len(response.Results))
	for i, result := range response.Results {
		fmt.Printf("%d. %s\n   URL: %s\n   Description: %s\n\n", i+1, result.Title, result.URL, result.Description)
	}

	// Get next page results
	fmt.Println("Getting page 2 results...")
	nextResponse, err := client.Search(ctx, params.NextPage())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nFound %d results in page 2:\n\n", len(nextResponse.Results))
	for i, result := range nextResponse.Results {
		fmt.Printf("%d. %s\n   URL: %s\n   Description: %s\n\n", i+1, result.Title, result.URL, result.Description)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Example 1: Basic news search
	fmt.Println("\n=== Basic News Search ===")
	news, err := client.News(ctx, &ddgsearch.NewsParams{
		Query:      "artificial intelligence",
		Region:     ddgsearch.RegionUS,
		SafeSearch: ddgsearch.SafeSearchModerate,
		MaxResults: 5,
	})
	if err != nil {
		log.Printf("Basic search error: %v", err)
	} else {
		printNewsResults("Basic Search", news)
	}

	// Example 2: News search with time range
	fmt.Println("\n=== News Search with Time Range ===")
	newsWithTime, err := client.News(ctx, &ddgsearch.NewsParams{
		Query:      "technology",
		Region:     ddgsearch.RegionUS,
		SafeSearch: ddgsearch.SafeSearchModerate,
		TimeRange:  ddgsearch.TimeRangeDay,
		MaxResults: 3,
	})
	if err != nil {
		log.Printf("Time range search error: %v", err)
	} else {
		printNewsResults("Time Range Search", newsWithTime)
	}

	// Example 3: News search with different region
	fmt.Println("\n=== News Search with Different Region ===")
	newsJP, err := client.News(ctx, &ddgsearch.NewsParams{
		Query:      "tech news",
		Region:     ddgsearch.RegionJP,
		SafeSearch: ddgsearch.SafeSearchModerate,
		MaxResults: 3,
	})
	if err != nil {
		log.Printf("Japanese region search error: %v", err)
	} else {
		printNewsResults("Japanese Region Search", newsJP)
	}
}

func printNewsResults(title string, news *ddgsearch.NewsResponse) {
	fmt.Printf("\n%s Results:\n", title)
	fmt.Printf("Found %d results\n", len(news.Results))
	for i, result := range news.Results {
		fmt.Printf("\n%d. %s\n", i+1, result.Title)
		fmt.Printf("   Source: %s\n", result.Source)
		fmt.Printf("   Date: %s\n", result.Date)
		fmt.Printf("   URL: %s\n", result.URL)
		if result.Image != "" {
			fmt.Printf("   Image: %s\n", result.Image)
		}
		fmt.Printf("   Summary: %s\n", result.Body)
	}
}
