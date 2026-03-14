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

func (manager *AppointmentManager) Add(ctx context.Context, payload *domain.Appointment) (*domain.Appointment, *core.Exception) {
	existentAppointment, err := existsByStartAndEndTime(ctx, manager, payload.StartTime, payload.EndTime)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error when verify availability"))
	}

	if existentAppointment {
		return nil, core.Confilct(core.WithMessage("could not schedule appointment"))
	}

	appointment, err := manager.repository.Add(ctx, payload)
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

	manager.emailProvider.Send(message)

	return appointment, nil
}

func (manager *AppointmentManager) GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, *core.Exception) {
	appointments, err := manager.repository.GetAllAppointmentsByProfessionalId(ctx, professionalId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get all appointment by professional id"), core.WithError(err))
	}
	return appointments, nil
}

func (manager *AppointmentManager) GetAppointmentById(ctx context.Context, appointmentId string) (*domain.Appointment, *core.Exception) {
	appointment, err := manager.repository.GetAppointmentById(ctx, appointmentId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return appointment, nil
}

func (manager *AppointmentManager) DeleteAppointment(ctx context.Context, appointmentId string) *core.Exception {
	err := manager.repository.DeleteAppointment(ctx, appointmentId)
	if err != nil {
		return core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return nil
}

func existsByStartAndEndTime(ctx context.Context, manager *AppointmentManager, startTime, endTime string) (bool, error) {
	exist, err := manager.repository.ExistsByStartAndEndTime(ctx, startTime, endTime)
	if err != nil {
		return false, err
	}

	return exist != false, nil
}
