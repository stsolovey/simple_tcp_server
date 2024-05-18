package server

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/stsolovey/simple_tcp_server/internal/config"
)

type Server struct {
	cfg      *config.Config
	log      *logrus.Logger
	listener net.Listener
}

func CreateServer(config *config.Config, logger *logrus.Logger) *Server {
	return &Server{
		cfg: config,
		log: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {

	var err error

	s.listener, err = net.Listen("tcp", s.cfg.Host+":"+s.cfg.Port)
	if err != nil {
		return fmt.Errorf("Error listening on port: %w", err)
	}

	defer s.listener.Close()
	s.log.Info("Server is listening on " + s.cfg.Host + ":" + s.cfg.Port)

	go func() {
		<-ctx.Done()
		s.listener.Close()
	}()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			switch {
			case ctx.Err() != nil:
				s.log.Info("Server stopped")
				return nil
			case strings.Contains(err.Error(), "use of closed network connection"):
				return nil
			default:
				s.log.Errorf("Error accepting connection: %s", err)
			}

			continue
		}

		go handleConnection(conn, s.log)
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

func (s *Server) Close() {
	s.listener.Close()
}
