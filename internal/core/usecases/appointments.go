package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/emailtemplates"

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
		manager.logger.Error("Error checking schedule availability.", "use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error when verify availability"))
	}

	if existentAppointment {
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
