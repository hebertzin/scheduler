package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Broker struct {
	broker       *amqp.Channel
	routingKey   string
	exchangeName string
}

func NewPublisher(routingKey, exchangeName string) *Broker {
	return &Broker{
		routingKey:   routingKey,
		exchangeName: exchangeName,
	}
}

func (b *Broker) Publish(ctx context.Context, event outbound.Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal event payload: %w", err)
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal envelope: %w", err)
	}

	if err := b.broker.PublishWithContext(
		ctx,
		b.exchangeName,
		b.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	); err != nil {
		return fmt.Errorf("rabbitmq publish: %w", err)
	}

	return nil
}
