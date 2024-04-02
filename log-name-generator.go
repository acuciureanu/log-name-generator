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

	// Generate log filenames
	dateFormats := []string{"2006-01-02", "20060102", "2006_01_02", "2006-Jan-02", "200601021504"}
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
	fmt.Fprintf(flag.CommandLine.Output(), "This program generates log filenames for a specified number of past days based on provided seed names or defaults.\n\n")
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nExamples:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  %s -days 7 -ext txt -log-names 'server,error'\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  This command generates server and error log filenames for the past 7 days with a .txt extension.\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  %s -days 5 -seed-file seeds.txt\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  This command uses seed names from 'seeds.txt' to generate log filenames for the past 5 days.\n")
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
