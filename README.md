# makedatelist

A command-line tool that generates a list of dates between two specified dates. Available in both Python and Go implementations.

## Features

- **Flexible date formats**: Supports multiple date input formats (YYYY-MM-DD, YYYYMMDD, MM/DD/YYYY, etc.)
- **Customizable output format**: Uses strftime-style format strings
- **Flexible argument order**: Flags can be placed anywhere in the command line
- **Cross-platform**: Available for Linux, Windows, and macOS

## Installation

### Go Version (Recommended)

```bash
# Clone the repository
git clone <repository-url>
cd makedatelist

# Build the binary
make build

# Install to system (optional)
make install
```

### Python Version

```bash
# Make the script executable
chmod +x makedatelist.py

# Run directly
./makedatelist.py
```

## Usage

### Basic Usage

```bash
# Default date range (2016-01-01 to 2016-01-31)
makedatelist

# Specify start and end dates
makedatelist 2024-01-01 2024-01-03

# With custom format
makedatelist 2024-01-01 2024-01-03 -f "%Y/%m/%d"
```

### Supported Date Formats

The tool accepts various date formats:

- `YYYY-MM-DD` (2024-01-01)
- `YYYYMMDD` (20240101)
- `MM/DD/YYYY` (01/01/2024)
- `DD/MM/YYYY` (01/01/2024)
- `Jan 1, 2024`
- `January 1, 2024`
- And more...

### Output Formats

Use strftime-style format strings:

- `%Y-%m-%d` → 2024-01-01 (default)
- `%Y/%m/%d` → 2024/01/01
- `%B %d, %Y` → January 01, 2024
- `%Y%m%d` → 20240101
- `%A, %B %d, %Y` → Monday, January 01, 2024

### Examples

```bash
# Basic usage
makedatelist 2024-01-01 2024-01-03
# Output:
# 2024-01-01
# 2024-01-02

# Custom format
makedatelist -f "%Y/%m/%d" 2024-01-01 2024-01-03
# Output:
# 2024/01/01
# 2024/01/02

# Month name format
makedatelist 2024-01-01 2024-01-03 -f "%B %d, %Y"
# Output:
# January 01, 2024
# January 02, 2024

# Different date formats
makedatelist "01/01/2024" "01/03/2024"
makedatelist "Jan 1, 2024" "Jan 3, 2024"
```

## Command Line Options

```
Usage: makedatelist [start_date] [end_date] [options]
       makedatelist [options] [start_date] [end_date]

Arguments:
  start_date  start date (supports multiple formats)
  end_date    end date (supports multiple formats)

Options:
  -f, --format string
        format - default %Y-%m-%d (supports strftime-style format)
  -h, --help
        show this help message
```

## Development

### Building

```bash
# Development build
make dev

# Optimized build
make build

# Cross-compile for all platforms
make build-all

# Run tests
make test

# Clean build artifacts
make clean
```

### Available Make Targets

- `build`: Create optimized binary
- `dev`: Create development binary with debug info
- `test`: Run functionality tests
- `clean`: Remove build artifacts
- `install`: Install to system (requires sudo)
- `build-all`: Cross-compile for Linux, Windows, and macOS
