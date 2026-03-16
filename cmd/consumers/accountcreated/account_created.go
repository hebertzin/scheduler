package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/infra/emailtemplates"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type AccountCreatedConsumer struct {
	conn        *amqp.Connection
	channel     *amqp.Channel
	logger      *logrus.Logger
	queueName   string
	emailSender outbound.EmailSender
}

func NewAccountCreatedConsumer(amqpURL, queueName string, sender outbound.EmailSender, logger *logrus.Logger) (*AccountCreatedConsumer, error) {
	conn, err := amqp.Dial(amqpURL)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	err = ch.Qos(
		10,
		0,
		false,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &AccountCreatedConsumer{
		conn:        conn,
		channel:     ch,
		queueName:   queueName,
		emailSender: sender,
		logger:      logger,
	}, nil
}

func (c *AccountCreatedConsumer) Consume(ctx context.Context) error {
	deliveries, err := c.channel.Consume(
		c.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.logger.Println("consumer started, waiting messages from queue=%s", c.queueName)

	for {
		select {
		case <-ctx.Done():
			c.logger.Println("consumer shutting down")
			return nil

		case delivery, ok := <-deliveries:
			if !ok {
				return fmt.Errorf("deliveries channel closed")
			}

			if err := c.handleMessage(delivery); err != nil {
				c.logger.Println("failed to process message: %v", err)

				if nackErr := delivery.Nack(false, true); nackErr != nil {
					c.logger.Println("failed to nack message: %v", nackErr)
				}
				continue
			}

			if err := delivery.Ack(false); err != nil {
				c.logger.Println("failed to ack message: %v", err)
			}
		}
	}
}

func (c *AccountCreatedConsumer) handleMessage(delivery amqp.Delivery) error {
	var event outbound.Event
	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		return fmt.Errorf("invalid event payload: %w", err)
	}

	var payload inbound.AccountCreatedEvent
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		return fmt.Errorf("invalid account created payload: %w", err)
	}

	if payload.Email == "" {
		return fmt.Errorf("payload without email")
	}

	data := emailtemplates.AccountCreatedData{
		Email: payload.Email,
	}

	body, err := emailtemplates.RenderAccountCreated(data)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	message := domain.EmailMessage{
		To:      []string{payload.Email},
		Subject: "Account created",
		Message: body,
	}

	c.emailSender.Send(message)

	return nil
}
