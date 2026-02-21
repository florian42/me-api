.PHONY: all build run test clean fmt vet tidy dev

BINARY_NAME=me-api
BUILD_DIR=bin

all: build

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

run:
	go run .

dev:
	@which air > /dev/null 2>&1 || (echo "Installing air..."; go install github.com/air-verse/air@latest)
	air

test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned build artifacts"

lint: fmt vet
	@echo "Running linters..."
