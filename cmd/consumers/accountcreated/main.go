package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hebertzin/scheduler/internal/infra/config/env"

	"github.com/hebertzin/scheduler/internal/infra/emailprovider"
	"github.com/sirupsen/logrus"
)

func main() {
	_, err := env.LoadConfiguration("/configs/config.json")
	if err != nil {
		panic(err)
	}

	logger := logrus.New()

	emailSender := emailprovider.NewSMPT(
		"",
		"",
		"",
	)

	consumer, err := NewAccountCreatedConsumer(
		"",
		"account.created",
		emailSender,
		logger,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := consumer.Consume(ctx); err != nil {
			logger.WithError(err).Error("consumer stopped with error")
			cancel()
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	logger.Info("shutting down consumer")
	cancel()
}
