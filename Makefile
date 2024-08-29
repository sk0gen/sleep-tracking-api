# Simple Makefile for a Go project
-include .env

# Build the application
all: build

build:
	@echo "Building..."

	@go build -o main main.go

# Run the application
run:
	@go run main.go serve;

# Create DB container
docker-up:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

new_migration:
	docker run --rm -v $(PWD)/internal/database/migrations:/migrations --network host migrate/migrate create -ext sql -dir /migrations -seq $(name); \


migration_up:
	@go run main.go migrate

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

sqlc:
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc generate

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

proto:
	rm -f internal/pb/*.go
	mkdir -p internal/pb
	protoc --proto_path=internal/proto --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative internal/proto/*.proto

swagger:
	cd internal/api && swag init -g server.go --parseDependency
	cd ../..

.PHONY: all build run docker-up new_migration migration_up docker-down sqlc test clean