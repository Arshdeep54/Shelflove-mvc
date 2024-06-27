GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
AIR := $(GOPATH_BIN)/air
GOLANGCI_LINT := $(GOPATH_BIN)/golangci-lint
GO_PACKAGES := $(shell go list ./... | grep -v vendor)
BUILD_OUTPUT := ./target/shelflove
BUILD_INPUT := cmd/main.go
UNAME := $(shell uname)


build:
	@echo "Building..."
	@test -d target || mkdir target
	@$(GO) build -o $(BUILD_OUTPUT) $(BUILD_INPUT)
	@echo "Built as $(BUILD_OUTPUT)"
dev:
	@echo "Starting development server..."
	@$(AIR)