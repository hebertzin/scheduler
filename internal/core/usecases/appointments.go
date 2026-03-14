package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/infra/emailtemplates"

	"github.com/sirupsen/logrus"
)

type AppointmentManager struct {
	repository          outbound.AppointmentRepository
	availabilityManager *AvailabilityManager
	emailProvider       outbound.EmailSender
	logger              *logrus.Logger
}

func NewAppointment(
	repository outbound.AppointmentRepository,
	availabilityManager *AvailabilityManager,
	logger *logrus.Logger,
	emailProvider outbound.EmailSender,
) inbound.AppointmentUseCase {
	return &AppointmentManager{
		repository:          repository,
		availabilityManager: availabilityManager,
		emailProvider:       emailProvider,
		logger:              logger,
	}
}

func (manager *AppointmentManager) Add(ctx context.Context, payload *domain.Appointment) (*domain.Appointment, *core.Exception) {
	exists, err := manager.availabilityManager.ExistByStartAndEndTime(ctx, payload.StartTime, payload.EndTime)
	if err != nil {
		manager.logger.Error("Error checking schedule availability.", "use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error when verify availability"))
	}

	if exists {
		manager.logger.Error("Time slot not available, time slot already scheduled.", "use_case_manager", "err", err.Error())

		return nil, core.Confilct(core.WithMessage("It was not possible to schedule."))
	}

	appointment, err := manager.repository.Add(ctx, payload)
	if err != nil {
		manager.logger.Error("An error occurred while scheduling", "use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error creating appointment"), core.WithError(err))
	}

	appointmentConfirmation := emailtemplates.AppointmentConfirmationData{
		Name:      appointment.Email,
		StartTime: appointment.StartTime,
		EndTime:   appointment.EndTime,
	}

	body, _ := emailtemplates.RenderAppointmentConfirmation(appointmentConfirmation)

	message := domain.EmailMessage{
		From:    "hebertsantosdeveloper@gmail.com",
		To:      []string{appointment.Email},
		Subject: emailtemplates.AppointmentConfirmationSubject,
		Message: body,
	}

	manager.emailProvider.Send(message)

	manager.logger.Println("Appointment created and confirmation email was sent")
	return appointment, nil
}

func (manager *AppointmentManager) GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, *core.Exception) {
	appointments, err := manager.repository.GetAllAppointmentsByProfessionalId(ctx, professionalId)
	if err != nil {
		manager.logger.Error("Error get all appointments by professional owner", "use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error get all appointment by professional id"), core.WithError(err))
	}

	return appointments, nil
}

func (manager *AppointmentManager) GetAppointmentById(ctx context.Context, appointmentId string) (*domain.Appointment, *core.Exception) {
	appointment, err := manager.repository.GetAppointmentById(ctx, appointmentId)
	if err != nil {
		manager.logger.Error("Error get appointment detail", "use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}

	return appointment, nil
}

func (manager *AppointmentManager) DeleteAppointment(ctx context.Context, appointmentId string) *core.Exception {
	err := manager.repository.DeleteAppointment(ctx, appointmentId)
	if err != nil {
		manager.logger.Error("Error deleting appointment", "use_case_manager", "err", err.Error())

		return core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return nil
}
