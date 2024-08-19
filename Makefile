# Simple Makefile for a Go project
-include .env.test
-include .env

# Build the application
all: build

build:
	@echo "Building..."

	@go build -o main cmd/main.go

# Run the application
run:
	@go run cmd/main.go

# Create DB container
docker-up:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Migrate DB
migration_up:
	migrate -path internal/database/migrations/ -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@localhost:5432/$(DB_DATABASE)?sslmode=disable" -verbose up

migration_down:
	migrate -path internal/database/migrations/ -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@localhost:5432/$(DB_DATABASE)?sslmode=disable" -verbose down
# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

sqlc:
	sqlc generate

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

.PHONY: all build run docker-up migration_up docker-down sqlc test clean