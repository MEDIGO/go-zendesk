include .env

test:
	@echo "Running tests..."
	@go test ./zendesk
.PHONY: test

test-integration:
	@echo "Running integration tests..."
	@go test ./test/integration
.PHONY: test-integration
