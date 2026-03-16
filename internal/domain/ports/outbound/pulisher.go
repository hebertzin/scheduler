package outbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
)

type Publisher interface {
	Publish(ctx context.Context, event inbound.Event) error
}
