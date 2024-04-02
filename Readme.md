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
  log-name-generator -days 5 -seed-file seeds.txt
  This command uses seed names from 'seeds.txt' to generate log filenames for the past 5 days.
```

## License

This tool is free to use under the MIT License.
