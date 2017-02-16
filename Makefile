include .env
export $(shell sed 's/=.*//' .env)

all: lint test
.PHONY: all

lint:
	@echo "Linting source code..."
	@go vet ./...
.PHONY: lint

test:
	@echo "Running tests..."
	@go test -v ./...
.PHONY: test
