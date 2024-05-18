package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stsolovey/simple_tcp_server/internal/config"
	"github.com/stsolovey/simple_tcp_server/internal/server"
)

func main() {
	log := logrus.New()

	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Panic("Error loading .env file")
	}

	cfg, err := config.New()
	if err != nil {
		log.WithError(err).Panic("Failed to load coniguration")
	}

	tcpServer := server.CreateServer(cfg, log)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		log.Info("Shutting down server")
		tcpServer.Close()
	}()

	err = tcpServer.Start(ctx)
	if err != nil {
		log.WithError(err).Panic("Failed to start server")
	}

	<-ctx.Done()
}
