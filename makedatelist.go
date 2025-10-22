package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// DateRange represents a range of dates
type DateRange struct {
	Start time.Time
	End   time.Time
}

// NewDateRange creates a new DateRange from start and end date strings
func NewDateRange(startStr, endStr string) (*DateRange, error) {
	start, err := ParseDate(startStr)
	if err != nil {
		return nil, err
	}

	end, err := ParseDate(endStr)
	if err != nil {
		return nil, err
	}

	return &DateRange{Start: start, End: end}, nil
}

// GenerateDates generates dates in the range (excluding end date)
func (dr *DateRange) GenerateDates() []time.Time {
	var dates []time.Time
	for d := dr.Start; d.Before(dr.End); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}
	return dates
}

// ParseDate attempts to parse a date string in various common formats
func ParseDate(dateStr string) (time.Time, error) {
	// List of common date formats to try
	formats := []string{
		"20060102",        // YYYYMMDD
		"2006-01-02",      // YYYY-MM-DD
		"2006/01/02",      // YYYY/MM/DD
		"01/02/2006",      // MM/DD/YYYY
		"02/01/2006",      // DD/MM/YYYY
		"02-01-2006",      // DD-MM-YYYY
		"01-02-2006",      // MM-DD-YYYY
		"Jan 2, 2006",     // Mon Jan 2, 2006
		"January 2, 2006", // January 2, 2006
		"2006-1-2",        // YYYY-M-D
		"2006/1/2",        // YYYY/M/D
		"1/2/2006",        // M/D/YYYY
		"2/1/2006",        // D/M/YYYY
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("incorrect timestamp: %s", dateStr)
}

// ConvertFormat converts strftime-style format to Go time format
func ConvertFormat(format string) string {
	// Common strftime to Go time format mappings
	replacements := map[string]string{
		"%Y": "2006",                     // 4-digit year
		"%y": "06",                       // 2-digit year
		"%m": "01",                       // month (01-12)
		"%d": "02",                       // day (01-31)
		"%H": "15",                       // hour (00-23)
		"%M": "04",                       // minute (00-59)
		"%S": "05",                       // second (00-59)
		"%B": "January",                  // full month name
		"%b": "Jan",                      // abbreviated month name
		"%A": "Monday",                   // full weekday name
		"%a": "Mon",                      // abbreviated weekday name
		"%j": "002",                      // day of year (001-366)
		"%U": "00",                       // week number of year (00-53)
		"%W": "00",                       // week number of year (00-53)
		"%w": "0",                        // weekday (0-6, Sunday is 0)
		"%x": "01/02/06",                 // date representation
		"%X": "15:04:05",                 // time representation
		"%c": "Mon Jan 02 15:04:05 2006", // date and time
		"%Z": "MST",                      // timezone name
		"%z": "-0700",                    // timezone offset
	}

	result := format
	for strftime, goFormat := range replacements {
		result = strings.ReplaceAll(result, strftime, goFormat)
	}

	return result
}

// parseArguments manually parses command line arguments to support flags anywhere
func parseArguments() (startDate, endDate, format string) {
	// Set default values
	format = "%Y-%m-%d"
	startDate = "2016-01-01"
	endDate = "2016-01-31"

	args := os.Args[1:] // skip program name
	var positionalArgs []string

	// Parse arguments manually
	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == "-f" || arg == "--format":
			// Handle format flag with space
			if i+1 < len(args) {
				format = args[i+1]
				i++ // skip the format value
			} else {
				fmt.Fprintf(os.Stderr, "Error: %s flag requires a value\n", arg)
				os.Exit(1)
			}
		case strings.HasPrefix(arg, "-f=") || strings.HasPrefix(arg, "--format="):
			// Handle format flag with equals
			if strings.HasPrefix(arg, "-f=") {
				format = strings.TrimPrefix(arg, "-f=")
			} else {
				format = strings.TrimPrefix(arg, "--format=")
			}
		case arg == "-h" || arg == "--help":
			// Show help and exit
			showHelp()
			os.Exit(0)
		case strings.HasPrefix(arg, "-"):
			// Unknown flag
			fmt.Fprintf(os.Stderr, "Error: unknown flag %s\n", arg)
			os.Exit(1)
		default:
			// Positional argument
			positionalArgs = append(positionalArgs, arg)
		}
	}

	// Set dates based on positional arguments
	switch len(positionalArgs) {
	case 0:
		// Use defaults
	case 1:
		startDate = positionalArgs[0]
		// endDate remains default
	default:
		startDate = positionalArgs[0]
		endDate = positionalArgs[1]
	}

	return startDate, endDate, format
}

// showHelp displays the help message
func showHelp() {
	// Get program name without path
	programName := "makedatelist"
	if len(os.Args) > 0 {
		parts := strings.Split(os.Args[0], "/")
		programName = parts[len(parts)-1]
	}

	fmt.Fprintf(os.Stderr, "This script make a list of range of specific two dates.\n\n")
	fmt.Fprintf(os.Stderr, "Create Date: 2016-07-25\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [start_date] [end_date] [options]\n", programName)
	fmt.Fprintf(os.Stderr, "       %s [options] [start_date] [end_date]\n", programName)
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Arguments:\n")
	fmt.Fprintf(os.Stderr, "  start_date  start date (supports multiple formats: YYYY-MM-DD, YYYYMMDD, MM/DD/YYYY, etc.)\n")
	fmt.Fprintf(os.Stderr, "  end_date    end date (supports multiple formats: YYYY-MM-DD, YYYYMMDD, MM/DD/YYYY, etc.)\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	fmt.Fprintf(os.Stderr, "  -f, --format string\n")
	fmt.Fprintf(os.Stderr, "        format - default %%Y-%%m-%%d (supports strftime-style format)\n")
	fmt.Fprintf(os.Stderr, "  -h, --help\n")
	fmt.Fprintf(os.Stderr, "        show this help message\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Examples:\n")
	fmt.Fprintf(os.Stderr, "  %s\n", programName)
	fmt.Fprintf(os.Stderr, "  %s 2024-01-01 2024-01-03\n", programName)
	fmt.Fprintf(os.Stderr, "  %s -f \"%%Y/%%m/%%d\" 2024-01-01 2024-01-03\n", programName)
	fmt.Fprintf(os.Stderr, "  %s 2024-01-01 2024-01-03 -f \"%%B %%d, %%Y\"\n", programName)
}

func main() {
	// Parse command line arguments
	startDate, endDate, format := parseArguments()

	// Create date range
	dateRange, err := NewDateRange(startDate, endDate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	// Generate and print dates
	dates := dateRange.GenerateDates()
	goFormat := ConvertFormat(format)

	for _, date := range dates {
		fmt.Println(date.Format(goFormat))
	}
}
