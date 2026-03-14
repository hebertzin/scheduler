package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"

	"github.com/sirupsen/logrus"
)

type AppointmentManager struct {
	repository    domain.AppointmentRepository
	emailProvider domain.EmailSender
	logger        *logrus.Logger
}

func NewAppointment(repository domain.AppointmentRepository, logger *logrus.Logger, emailProvider domain.EmailSender) domain.AppointmentUseCase {
	return &AppointmentManager{
		repository:    repository,
		emailProvider: emailProvider,
		logger:        logger,
	}
}

func (s *AppointmentManager) Add(ctx context.Context, appointment *domain.Appointment) (*domain.Appointment, *core.Exception) {
	appointment, err := s.repository.Add(ctx, appointment)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating appointment"), core.WithError(err))
	}

	message := domain.EmailMessage{
		From:    "hebertsantosdeveloper@gmail.com",
		To:      []string{appointment.Email},
		Subject: "Appointment Confirmation",
		Message: `Hello, Your appointment has been successfully scheduled.
                 If you need to reschedule or cancel, please reply to this email or contact our support team.
                Thank you.`,
	}

	s.emailProvider.Send(message)

	return appointment, nil
}

func (s *AppointmentManager) GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, *core.Exception) {
	appointments, err := s.repository.GetAllAppointmentsByProfessionalId(ctx, professionalId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get all appointment by professional id"), core.WithError(err))
	}
	return appointments, nil
}

func (s *AppointmentManager) GetAppointmentById(ctx context.Context, appointmentId string) (*domain.Appointment, *core.Exception) {
	appointment, err := s.repository.GetAppointmentById(ctx, appointmentId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return appointment, nil
}

func (s *AppointmentManager) DeleteAppointment(ctx context.Context, appointmentId string) *core.Exception {
	err := s.repository.DeleteAppointment(ctx, appointmentId)
	if err != nil {
		return core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return nil
}
