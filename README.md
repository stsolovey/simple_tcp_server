```markdown
# Simple TCP Server

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This is a simple TCP server written in Go that listens on `localhost:4080`. It logs incoming messages and sends back a confirmation message.

## Getting Started

### Prerequisites

- Go (https://golang.org/dl/)

### Installing

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/simple_tcp_server.git
   cd simple_tcp_server
   ```

2. Install dependencies:

   ```bash
   go get ./...
   go mod tidy
   ```

### Running the Server

You can run the server using the following Makefile command:

```bash
make up
```

## Usage

To interact with the server, you can use a tool like `telnet` or `nc` (netcat):

```bash
telnet localhost 4080
# or
nc localhost 4080
```

Type a message and press Enter. The server will log your message and send a confirmation back.

```
