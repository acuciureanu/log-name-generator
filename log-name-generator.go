package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func printUsage() {
	fmt.Println("Usage: log-generator <number_of_days>")
	fmt.Println("Generate log filenames for the specified number of days in the past, up to the current day.")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	days, err := strconv.Atoi(os.Args[1])
	if err != nil || days < 1 {
		fmt.Println("Error: The number of days should be a positive integer.")
		printUsage()
		os.Exit(1)
	}

	logTypes := []string{
		"access", "access_log", "authorizenet", "development",
		"error", "error_log", "exception", "librepag",
		"log", "old", "payment", "payment_authorizenet",
		"payment_paypal_express", "production", "server", "test", "www-error",
	}

	dateFormats := []string{
		"2006-01-02",   // ISO 8601: YYYY-MM-DD
		"20060102",     // Compact: YYYYMMDD
		"2006_01_02",   // Underscore: YYYY_MM_DD
		"2006-Jan-02",  // Hyphenated Short Month: YYYY-Mon-DD
		"200601021504", // Compact Timestamp: YYYYMMDDHHMM
	}

	for day := 0; day < days; day++ {
		timestamp := time.Now().AddDate(0, 0, -day)
		for _, logType := range logTypes {
			for _, format := range dateFormats {
				filename := fmt.Sprintf("%s-%s.log", logType, timestamp.Format(format))
				fmt.Println(filename)
			}
		}
	}
}
