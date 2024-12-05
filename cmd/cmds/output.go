package cmds

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/kuhahalong/ddgsearch"
)

// printResults 根据输出类型打印搜索结果
func printResults(results interface{}) {
	switch outputType {
	case "json":
		printJSON(results)
	case "text":
		printText(results)
	case "csv":
		printCSV(results)
	default:
		fmt.Fprintf(os.Stderr, "Unsupported output type: %s\n", outputType)
	}
}

// printJSON 以 JSON 格式输出结果
func printJSON(results interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
	}
}

// printText 以文本格式输出结果
func printText(results interface{}) {
	switch v := results.(type) {
	case *ddgsearch.SearchResponse:
		for _, result := range v.Results {
			fmt.Printf("Title: %s\n", result.Title)
			fmt.Printf("URL: %s\n", result.URL)
			fmt.Printf("Description: %s\n", result.Description)
			fmt.Println(strings.Repeat("-", 80))
		}
	case *ddgsearch.NewsResponse:
		for _, result := range v.Results {
			fmt.Printf("Title: %s\n", result.Title)
			fmt.Printf("URL: %s\n", result.URL)
			fmt.Printf("Source: %s\n", result.Source)
			fmt.Printf("Date: %s\n", result.Date)
			fmt.Println(strings.Repeat("-", 80))
		}
	}
}

// printCSV 以 CSV 格式输出结果
func printCSV(results interface{}) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	switch v := results.(type) {
	case *ddgsearch.SearchResponse:
		writer.Write([]string{"Title", "URL", "Description"})
		for _, result := range v.Results {
			writer.Write([]string{result.Title, result.URL, result.Description})
		}
	case *ddgsearch.NewsResponse:
		writer.Write([]string{"Title", "URL", "Source", "Date"})
		for _, result := range v.Results {
			writer.Write([]string{result.Title, result.URL, result.Source, result.Date})
		}
	}
}

// saveResults 将结果保存到文件
func saveResults(results interface{}) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	ext := strings.ToLower(strings.TrimPrefix(outputFile[strings.LastIndex(outputFile, "."):], "."))
	switch ext {
	case "json":
		return saveJSON(file, results)
	case "csv":
		return saveCSV(file, results)
	default:
		return fmt.Errorf("unsupported output format: %s", ext)
	}
}

// saveJSON 将结果以 JSON 格式保存到文件
func saveJSON(file *os.File, results interface{}) error {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}

// saveCSV 将结果以 CSV 格式保存到文件
func saveCSV(file *os.File, results interface{}) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	switch v := results.(type) {
	case *ddgsearch.SearchResponse:
		if err := writer.Write([]string{"Title", "URL", "Description"}); err != nil {
			return err
		}
		for _, result := range v.Results {
			if err := writer.Write([]string{result.Title, result.URL, result.Description}); err != nil {
				return err
			}
		}
	case *ddgsearch.NewsResponse:
		if err := writer.Write([]string{"Title", "URL", "Source", "Date"}); err != nil {
			return err
		}
		for _, result := range v.Results {
			if err := writer.Write([]string{result.Title, result.URL, result.Source, result.Date}); err != nil {
				return err
			}
		}
	}
	return nil
}
