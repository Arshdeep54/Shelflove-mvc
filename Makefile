.PHONY: build install help lint 

.DEFAULT_GOAL := help

GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
AIR := $(GOPATH_BIN)/air
GOLANGCI_LINT := ./bin/golangci-lint
GO_PACKAGES := $(shell go list ./... | grep -v vendor)
MIGRATE_PATH=./pkg/config/migrations/migration.go
BUILD_OUTPUT := ./target/shelflove
BUILD_INPUT := cmd/main.go
UNAME := $(shell uname)


help:
	@echo "Available targets:"
	@echo "  all      - Runs install ,migrate ,dev"
	@echo "  install  - Install dependencies and create vendor folder"
	@echo "  migrate  - Create a Database,tables and pushes some dummy data"
	@echo "  cleandb  - Clean the data in the tables"
	@echo "  build    - Build the project"
	@echo "  dev      - Start development server"
	@echo "  lint     - Run code linters "
	@echo "  test     - Run unit tests " 

all: install migrate dev
migrate:
	@$(GO) run ${MIGRATE_PATH}

cleandb:
	@CLEANDB=true $(GO) run ${MIGRATE_PATH}
	@CLEANDB=false 
lint:
	@$(GOLANGCI_LINT) run
	
test:
	@echo "tests yet to be created"
	
install:
	@echo "Installing dependencies..."
	@$(GO) mod download

vendor:
	@echo "Tidy up go.mod..."
	@$(GO) mod tidy
	@echo "Vendoring..."
	@$(GO) mod vendor
	@echo "Done!"

build:
	@echo "Building..."
	@test -d target || mkdir target
	@$(GO) build -o $(BUILD_OUTPUT) $(BUILD_INPUT)
	@echo "Built as $(BUILD_OUTPUT)"

dev:
	@echo "Starting development server..."
	@$(AIR)
