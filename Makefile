.PHONY: build install help lint 

.DEFAULT_GOAL := help

GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
AIR := $(GOPATH_BIN)/air
GOLANGCI_LINT := ./bin/golangci-lint
GO_PACKAGES := $(shell go list ./... | grep -v vendor)
MIGRATE_PATH=./pkg/config/migrations
BUILD_OUTPUT := shelflove
BUILD_INPUT := cmd/main.go
TEST_FOLDER:=./pkg/tests

help:
	@echo "Available targets:"
	@echo "  all      - Install ,erase all data, migrate,run "
	@echo "  install  - Install dependencies and create vendor folder"
	@echo "  migrate  - Create tables and pushes some dummy book data"
	@echo "  cleandb  - Clean the data in the tables"
	@echo "  build    - Build the project"
	@echo "  dev      - Start development server"
	@echo "  lint     - Run code linters "
	@echo "  test     - Run unit tests " 
	@echo "  host     - Host on Apache Server " 

all: install cleandb migrate dev
migrate:
	@$(GO) build -o ${MIGRATE_PATH}/migration ${MIGRATE_PATH}/migration.go
	@chmod +x ${MIGRATE_PATH}/migration
	@./${MIGRATE_PATH}/migration

cleandb:
	@$(GO) build -o ${MIGRATE_PATH}/migration ${MIGRATE_PATH}/migration.go
	@chmod +x ${MIGRATE_PATH}/migration
	@CLEANDB=true ./${MIGRATE_PATH}/migration

lint:
	@echo "Linting ..."
	@$(GOLANGCI_LINT) run
	
test:
	@echo "Testing"
	@go test ${TEST_FOLDER} -v
	
host:
	@echo "Hosting on apache server"
	@chmod +x ./host.sh
	@bash host.sh --sudo
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
