all: lint test
.PHONY: all

lint:
	@echo "Linting source code..."
	@go vet ./...
.PHONY: lint

test:
	@echo "Running tests..."
	@go test ./...
.PHONY: test
