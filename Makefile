# Makefile for makedatelist

.PHONY: build clean install test

# Build the binary
build:
	go build -ldflags="-s -w" -o makedatelist makedatelist.go

# Clean build artifacts
clean:
	rm -f makedatelist

# Install to system (requires sudo)
install: build
	sudo cp makedatelist /usr/local/bin/

# Run tests
test:
	go run makedatelist.go -h
	go run makedatelist.go "2024-01-01" "2024-01-03"
	go run makedatelist.go "2024-01-01" "2024-01-03" -f "%Y/%m/%d"

# Development build with debug info
dev:
	go build -o makedatelist makedatelist.go

# Cross-compile for different platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o makedatelist-linux-amd64 makedatelist.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o makedatelist-windows-amd64.exe makedatelist.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o makedatelist-darwin-amd64 makedatelist.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o makedatelist-darwin-arm64 makedatelist.go

# Default target
all: build
