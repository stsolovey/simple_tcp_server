package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	listener, err := net.Listen("tcp", "localhost:4080")
	if err != nil {
		log.WithError(err).Panic("Failed to initialize storage")
	}

	defer listener.Close()
	log.Info("Server is listening on localhost:4080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Errorf("Error accepting connection: %s", err)

			continue
		}

		go handleConnection(conn, log)
	}
}

func handleConnection(conn net.Conn, log *logrus.Logger) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Infof("Received: %s", text)

		_, err := fmt.Fprintf(conn, "You sent: %s\n", text)
		if err != nil {
			log.Errorf("Error writing to connection: %s", err)

			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Errorf("Error reading from connection: %s", err)
	}
}
