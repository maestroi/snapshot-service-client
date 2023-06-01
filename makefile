# Go Application Makefile

# Variables
BINARY_NAME := snapshot-service-client
MAIN_FILE := cmd/main.go

# Build the application
build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# Run the application locally
run:
	go run $(MAIN_FILE)

# Build the Docker image
docker-build:
	docker build -t snapshot-service-client .

# Run the Docker container
docker-run:
	docker run -it -v $(PWD):/app/data -f nimiq-v1 testnet /app/data

# Clean up generated files
clean:
	rm -f $(BINARY_NAME)

# Help command
help:
	@echo "Available commands:"
	@echo "- build: Build the application"
	@echo "- run: Run the application locally"
	@echo "- docker-build: Build the Docker image"
	@echo "- docker-run: Run the Docker container"
	@echo "- clean: Clean up generated files"
	@echo "- help: Show this help message"

.PHONY: build run docker-build docker-run clean help
