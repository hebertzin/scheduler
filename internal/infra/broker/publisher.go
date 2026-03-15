package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PublishingConfig struct {
	RoutingKey   string
	ExchangeName string
}

type Broker struct {
	ch              *amqp.Channel
	publisherConfig PublishingConfig
}

func NewPublisher(ch *amqp.Channel, publisherConfig PublishingConfig) *Broker {
	return &Broker{
		ch: ch,
		publisherConfig: PublishingConfig{
			RoutingKey:   publisherConfig.RoutingKey,
			ExchangeName: publisherConfig.ExchangeName,
		},
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

	if err := b.ch.PublishWithContext(
		ctx,
		b.publisherConfig.ExchangeName,
		b.publisherConfig.RoutingKey,
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
