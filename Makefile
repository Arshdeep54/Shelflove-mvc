GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
AIR := $(GOPATH_BIN)/air
GOLANGCI_LINT := ./bin/golangci-lint
GO_PACKAGES := $(shell go list ./... | grep -v vendor)
BUILD_OUTPUT := ./target/shelflove
BUILD_INPUT := cmd/main.go
UNAME := $(shell uname)


help:
	@echo "Available targets:"
	@echo "  install  - Install dependencies and create vendor folder"
	@echo "  build    - Build the project"
	@echo "  dev      - Start development server"
	@echo "  lint     - Run code linters "
	@echo "  test     - Run unit tests " 

migrate:
	@go run ./pkg/config/migrations/migration.go 

cleandb:
	@CLEANDB=true go run ./pkg/config/migrations/migration.go
	@CLEANDB=false 
lint:
	@$(GOLANGCI_LINT) run
	
install:
	@echo "Installing dependencies..."
	@go mod download
	@go mod vendor
build:
	@echo "Building..."
	@test -d target || mkdir target
	@$(GO) build -o $(BUILD_OUTPUT) $(BUILD_INPUT)
	@echo "Built as $(BUILD_OUTPUT)"

dev:
	@echo "Starting development server..."
	@$(AIR)
