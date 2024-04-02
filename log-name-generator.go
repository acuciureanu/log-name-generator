package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	// Set usage message
	flag.Usage = usage
	// Define and parse flags
	seedFile := flag.String("seed-file", "", "File containing the base name of log files")
	daysFlag := flag.Int("days", 0, "Number of days")
	ext := flag.String("ext", "log", "Extension for generated log files")
	logNames := flag.String("log-names", "", "Comma-separated base names of log files")
	dateFormat := flag.String("date-format", "", "Specify a single date format for log filenames (e.g., '20060102' for YYYYMMDD). See usage for more examples.")
	flag.Parse()

	// Check if number of days is provided
	if *daysFlag <= 0 {
		fmt.Println("Error: Number of days should be a positive integer")
		flag.Usage()
		os.Exit(1)
	}

	// Read seed from file if provided, otherwise use default or command line specified log names
	var seeds []string
	if *logNames != "" {
		seeds = parseLogNames(*logNames)
	} else {
		seeds = readSeedFile(*seedFile)
	}

	// Set dateFormats based on the provided dateFormat flag
	var dateFormats []string
	if *dateFormat != "" {
		dateFormats = []string{*dateFormat}
	} else {
		dateFormats = []string{"2006-01-02", "20060102", "2006_01_02", "2006-Jan-02", "200601021504"}
	}

	now := time.Now()

	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	for day := 0; day < *daysFlag; day++ {
		wg.Add(1) // Increment WaitGroup counter

		go func(day int) {
			defer wg.Done() // Decrement WaitGroup counter when goroutine finishes

			timestamp := now.AddDate(0, 0, -day)
			for _, seed := range seeds {
				for _, format := range dateFormats {
					fmt.Printf("%s-%s.%s\n", seed, timestamp.Format(format), *ext)
				}
			}
		}(day)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "This program generates log filenames for a specified number of past days based on provided seed names or defaults. It allows customization of the date format in the filenames.\n\n")
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nExamples:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  %s -days 7 -ext txt -log-names 'server,error'\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  This command generates server and error log filenames for the past 7 days with a .txt extension.\n")

	fmt.Fprintf(flag.CommandLine.Output(), "\nDate Format Examples:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  -date-format \"2006-01-02\" : Use ISO 8601 format (e.g., server-2024-04-02.log)\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  -date-format \"20060102\" : Use a compact format without separators (e.g., server-20240402.log)\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  -date-format \"2006_01_02\" : Use underscores as separators (e.g., server-2024_04_02.log)\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  -date-format \"2006-Jan-02\" : Include a three-letter month abbreviation (e.g., server-2024-Apr-02.log)\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  -date-format \"200601021504\" : Extend the compact date format to include hours and minutes (e.g., server-202404021530.log)\n")

	fmt.Fprintf(flag.CommandLine.Output(), "\nMore usage examples:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  %s -days 5 -seed-file seeds.txt\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  This command uses seed names from 'seeds.txt' to generate log filenames for the past 5 days.\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  %s -days 3 -date-format \"2006-Jan-02\"\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  This command generates log filenames for the past 3 days using the format with a three-letter month abbreviation.\n")
}

func parseLogNames(logNames string) []string {
	parts := strings.Split(logNames, ",")
	var parsedNames []string

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		trimmed = strings.Trim(trimmed, `"'`)

		if trimmed != "" {
			parsedNames = append(parsedNames, trimmed)
		}
	}

	return parsedNames
}

func readSeedFile(seedFile string) []string {
	if seedFile == "" {
		return getDefaultLogTypes()
	}

	file, err := os.Open(seedFile)
	if err != nil {
		fmt.Println("Error opening seed file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var seeds []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seeds = append(seeds, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading seed file:", err)
		os.Exit(1)
	}

	return seeds
}

func getDefaultLogTypes() []string {
	return []string{"access", "access_log", "authorizenet", "development", "error", "error_log", "exception", "librepag", "log", "old", "payment", "payment_authorizenet", "payment_paypal_express", "production", "server", "test", "www-error"}
}
