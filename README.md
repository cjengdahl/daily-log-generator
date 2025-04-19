# Daily Log Generator

## Overview

Generates a new markdown file named after the day of the month
within the nested directory path `<DAILY_LOG_DIRECTORY>/year/month`.
Intermediate directories will be created automatically (think `mkdir -p`).
The file is prefixed with a header with the day's date.

Example:

If the daily-log-generator were run on 1 January 1970,
a new file, and all directories in the file path, would be created
```
daily-logs
└── 1970
    └── 01
        └── 01.md
```

The file would contain a single line
```markdown
# Thursday, January 1st, 1970
```

## Build

Run the following to build `daily-log-generator` and to make it available for use.
- `go build -o dlg main.go`
- `mv dlg /usr/local/bin`

### Alternate

You can run `go install github.com/cjengdahl/daily-log-generator`
This will put the binary in `$GOBIN`.  So as long as `$GOBIN`
is in your path, you should be good to go.  If you plan to use
with launchd automation, you will need to modify your plist to reference
`$GOBIN`.

## Usage

Just run `dlg`, and check your home directory.

If you want a custom root directory path for your daily logs,
you can set the environment variable `DAILY_LOG_DIRECTORY`.
This must be an absolute directory path.

### Automation

This tool can be run manually, or you can automate it
with something like cron, or launchd to run daily.

### launchd (MacOS)

A plist is included in this repo for the convenience
of the user, if he/she so wishes to use the tool with launchd
on a Mac.  Directory customization is commented out, uncomment
as needed.

Simply copy the plist to the user launchd agents directory, and load the job.
- `cp daily-log-generator.plist ~/Library/LaunchAgents`
- `launchctl load -w ~/Library/LaunchAgents/daily-log-generator.plist`

To stop the job, unload and remove it
- `launchctl unload ~/Library/LaunchAgents/daily-log-generator.plist`
- `rm ~/Library/LaunchAgents/daily-log-generator.plist`


