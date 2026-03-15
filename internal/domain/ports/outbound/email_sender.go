package outbound

import "github.com/hebertzin/scheduler/internal/domain"

type EmailSender interface {
	Send(message domain.EmailMessage)
}
