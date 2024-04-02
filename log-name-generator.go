package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	// Define and parse flags
	seedFile := flag.String("seed-file", "", "File containing the base name of log files")
	daysFlag := flag.Int("days", 0, "Number of days")
	ext := flag.String("ext", "log", "Extension for generated log files")
	flag.Parse()

	// Check if number of days is provided
	if *daysFlag <= 0 {
		fmt.Println("Error: Number of days should be a positive integer")
		flag.Usage()
		os.Exit(1)
	}

	// Read seed from file if provided
	seeds := readSeedFile(*seedFile)

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
