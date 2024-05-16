up:
	go run main.go

tidy:
	gofumpt -w .
	gci write . --skip-generated -s standard -s default
	go mod tidy

lint: tidy
	golangci-lint run ./...

test:
	go test ./...

help:
	@echo "Available commands:"
	@echo "  up                - Start app"
	@echo "  tidy              - Format and tidy up the Go code"
	@echo "  lint              - Lint and format the project code"
	@echo "  test              - Start environment and run tests"