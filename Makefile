.PHONY: test

test:
	@echo "Running tests..."
	@go mod tidy
	@goimports -w .
	@go vet ./...
	@go test ./... -v