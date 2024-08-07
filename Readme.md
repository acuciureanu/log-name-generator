# Log Name Generator 🌈

Welcome! This tool makes naming log files a breeze, perfect for bug bounty hunters looking to generate fuzz lists effortlessly.

## Inspiration :v:

This tool was inspired from NahamSec's Youtube video [Don't Make This Recon Mistake // How To Bug Bounty](https://www.youtube.com/watch?v=YbIEXJhZxUk)

## What It Does

This tool makes log file names for any number of past days, up to today. It supports many log types and date formats.

### Log Types

We support many log types:

- 🚪 access
- 📜 access_log
- 💳 authorizenet
- 🔧 development
- ❗ error
- 📚 error_log
- 🚨 exception
- 💰 librepag
- 📝 log
- 🕰 old
- 💸 payment
- 💳 payment_authorizenet
- 🛒 payment_paypal_express
- 🏭 production
- 🖥 server
- 🧪 test
- 🌐 www-error

Example of a command which generates `.log` file names using only the default log types:

```bash
$ log-name-generator -date-format "2006-01-02" -days 1 -ext log
access-2024-07-08.txt
access_log-2024-07-08.txt
authorizenet-2024-07-08.txt
development-2024-07-08.txt
error-2024-07-08.txt
error_log-2024-07-08.txt
exception-2024-07-08.txt
librepag-2024-07-08.txt
log-2024-07-08.txt
old-2024-07-08.txt
payment-2024-07-08.txt
payment_authorizenet-2024-07-08.txt
payment_paypal_express-2024-07-08.txt
production-2024-07-08.txt
server-2024-07-08.txt
test-2024-07-08.txt
www-error-2024-07-08.txt
```

### Date Formats

It generates by default the following date formats:

- ISO 8601: YYYY-MM-DD
- Compact: YYYYMMDD
- Underscore: YYYY_MM_DD
- Hyphenated Short Month: YYYY-Mon-DD
- Compact Timestamp: YYYYMMDDHHMM

## How to Install

First, make sure you have Go on your computer. Then, run this command:

```bash
go install github.com/acuciureanu/log-name-generator@latest
```

## How to Use

To make log file names, run:

```bash
Usage of log-name-generator:
This program generates log filenames for a specified number of past days based on provided seed names or defaults.

  -date-format string
        Specify a single date format for log filenames (e.g., '20060102' for YYYYMMDD). See usage for more examples.
  -days int
        Number of days
  -ext string
        Extension for generated log files (default "log")
  -log-names string
        Comma-separated base names of log files
  -seed-file string
        File containing the base name of log files

Examples:
  log-name-generator -days 7 -ext txt -log-names 'server,error'
  This command generates server and error log filenames for the past 7 days with a .txt extension.

Date Format Examples:
  -date-format "2006-01-02" : Use ISO 8601 format (e.g., server-2024-04-02.log)
  -date-format "20060102" : Use a compact format without separators (e.g., server-20240402.log)
  -date-format "2006_01_02" : Use underscores as separators (e.g., server-2024_04_02.log)
  -date-format "2006-Jan-02" : Include a three-letter month abbreviation (e.g., server-2024-Apr-02.log)
  -date-format "200601021504" : Extend the compact date format to include hours and minutes (e.g., server-202404021530.log)

More usage examples:
  log-name-generator -days 5 -seed-file seeds.txt
  This command uses seed names from 'seeds.txt' to generate log filenames for the past 5 days.
  log-name-generator -days 3 -date-format "2006-Jan-02"
  This command generates log filenames for the past 3 days using the format with a three-letter month abbreviation.
```

#### Still not clear how to use the `-seed-file` argument ?

Let's say that `seeds.txt` is our seed file and it contains a few keywords which we want to include in the file names that we want to generate.

```txt
keyword1
keyword2
```

Now we can run `log-name-generator` with the `-seed-file` arguemnt and it'll generate file names with dates by including those keywords from our seed file.

```bash
$ log-name-generator -date-format "2006-01-02" -days 1 -ext txt -seed-file .\seed-file.txt
keyword1-2024-07-08.txt
keyword2-2024-07-08.txt
```

## License

This tool is free to use under the MIT License.
