# Log Name Generator 🌈

Welcome! This tool helps you name your log files easily. It's perfect for anyone who works with many logs and wants a simple way to name them.

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

You can choose from different date formats:

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
  -days int
        Number of days
  -ext string
        Extension for generated log files (default "log")
  -seed-file string
        File containing the base name of log files```

For example, for the last 7 days, run:

```bash
log-name-generator -days 7
```

## License

This tool is free to use under the MIT License.

Enjoy making your log files neat and organized! 🎉
