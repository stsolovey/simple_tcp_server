# Simple TCP Server

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This is a simple TCP server written in Go that listens on the host and port specified in the `.env` file. It logs incoming messages and sends back a confirmation message.

## Getting Started

### Prerequisites

- Go (https://golang.org/dl/)
- A `.env` file in the root directory with `HOST` and `PORT` specified. See `.env.example` for format.

### Installing

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/simple_tcp_server.git
   cd simple_tcp_server
   ```

2. Install dependencies:

   ```bash
   make tools       # Install required tools like gofumpt, gci, and golangci-lint
   go mod tidy      # Tidy up the dependencies
   ```

### Running the Server

You can run the server using the following Makefile command:

```bash
make up
```

This command will start the server as defined in `cmd/tcp_server/main.go`, loading configuration from the `.env` file.

## Usage

To interact with the server, you can use a tool like `telnet` or `nc` (netcat):

```bash
telnet [host] [port]
# or
nc [host] [port]
```

Replace `[host]` and `[port]` with the values specified in your `.env` file. Type a message and press Enter. The server will log your message and send a confirmation back.

## Development Commands

- **Start the server**: `make up`
- **Install necessary tools**: `make tools`
- **Format and tidy the code**: `make tidy`
- **Lint the code**: `make lint`
- **Run tests**: `make test`

## Contribution

Feel free to fork the repository, make improvements or customize as needed, and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.