include .env

test:
	@echo "Running tests..."
	@go test ./...
.PHONY: test
