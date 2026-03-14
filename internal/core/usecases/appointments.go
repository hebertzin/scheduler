package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/smtp"
	"github.com/sirupsen/logrus"
)

type AppointmentUseCase struct {
	repository domain.AppointmentRepository
	logger     *logrus.Logger
	smptConfig *smtp.SMPTConfig
}

func NewAppointmentUseCase(repository domain.AppointmentRepository, logger *logrus.Logger, smptConfig *smtp.SMPTConfig) domain.AppointmentUseCase {
	s := smtp.NewSMPT(smptConfig.Port, smptConfig.Password, smptConfig.Host)

	return &AppointmentUseCase{
		repository: repository,
		logger:     logger,
		smptConfig: s,
	}
}

func (s *AppointmentUseCase) Add(ctx context.Context, appointment *domain.Appointment) (*domain.Appointment, *core.Exception) {
	appointment, err := s.repository.Add(ctx, appointment)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating appointment"), core.WithError(err))
	}

	message := smtp.SMPTSendEmail{
		From:    "hebertsantosdeveloper@gmail.com",
		To:      []string{appointment.Email},
		Subject: "Appointment Confirmation",
		Message: `Hello, Your appointment has been successfully scheduled.
                 If you need to reschedule or cancel, please reply to this email or contact our support team.
                Thank you.`,
	}

	_ = s.smptConfig.Send(message)

	return appointment, nil
}

func (s *AppointmentUseCase) GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, *core.Exception) {
	appointments, err := s.repository.GetAllAppointmentsByProfessionalId(ctx, professionalId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get all appointment by professional id"), core.WithError(err))
	}
	return appointments, nil
}

func (s *AppointmentUseCase) GetAppointmentById(ctx context.Context, appointmentId string) (*domain.Appointment, *core.Exception) {
	appointment, err := s.repository.GetAppointmentById(ctx, appointmentId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return appointment, nil
}

func (s *AppointmentUseCase) DeleteAppointment(ctx context.Context, appointmentId string) *core.Exception {
	err := s.repository.DeleteAppointment(ctx, appointmentId)
	if err != nil {
		return core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return nil
}
