up:
	-go run cmd/tcp_server/main.go


tidy:
	gofumpt -w .
	gci write . --skip-generated -s standard -s default
	go mod tidy

lint: tidy
	golangci-lint run ./...

test:
	go test ./...

tools:
    go install mvdan.cc/gofumpt@latest
    go install github.com/daixiang0/gci@latest
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

help:
	@echo "Available commands:"
	@echo "  up                - Start app"
	@echo "  tidy              - Format and tidy up the Go code"
	@echo "  lint              - Lint and format the project code"
	@echo "  test              - Start environment and run tests"
	@echo "  tools             - Install necessary tools"