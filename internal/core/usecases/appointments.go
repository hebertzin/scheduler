package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/eventconstants"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/errors"
	"github.com/sirupsen/logrus"
)

type AppointmentManager struct {
	repository          outbound.AppointmentRepository
	availabilityManager Availability
	messaging           outbound.Publisher
	logger              *logrus.Logger
}

func NewAppointment(
	repository outbound.AppointmentRepository,
	availabilityManager Availability,
	messaging outbound.Publisher,
	logger *logrus.Logger,
) inbound.AppointmentUseCase {
	return &AppointmentManager{
		repository:          repository,
		availabilityManager: availabilityManager,
		messaging:           messaging,
		logger:              logger,
	}
}

func (manager *AppointmentManager) Add(ctx context.Context, payload *domain.Appointment) (*domain.Appointment, *errors.Exception) {
	exists, err := manager.availabilityManager.ExistByStartAndEndTime(ctx, payload.StartTime, payload.EndTime, payload.DayOfWeek)
	if err != nil {
		manager.logger.Error("Error checking schedule availability.", "use_case_manager", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error when verify availability"))
	}

	if exists {
		manager.logger.Error("Time slot not available, time slot already scheduled.", "use_case_manager")

		return nil, errors.Confilct(errors.WithMessage("It was not possible to schedule."))
	}

	appointment, err := manager.repository.Add(ctx, payload)
	if err != nil {
		manager.logger.Error("An error occurred while scheduling", "use_case_manager", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error creating appointment"), errors.WithError(err))
	}

	e := inbound.Event{
		Type:    eventconstants.AppointmentCreatedEventType,
		Payload: appointment,
	}

	err = manager.messaging.Publish(ctx, e)
	if err != nil {
		manager.logger.Println("Some error has been occurred publishing the message in broker.", "appointment_use_case_manager")

		// CONSISTENCY ISSUE: at this point the appointment was already persisted in the database
		// but the event failed to be published to the broker, leaving the system in an inconsistent state.
		// The consumer will never process this appointment, and retrying the request will result in a CONFLICT error.
		// TODO: implement the Outbox Pattern — persist the event in the same transaction as the appointment,
		// and use a background worker (relay) to publish it to the broker, ensuring at-least-once delivery.
		return nil, errors.Unexpected(errors.WithMessage("cannot publish the message in broker"))
	}

	manager.logger.Info("Appointment created and message was published")

	return appointment, nil
}

func (manager *AppointmentManager) GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, *errors.Exception) {
	appointments, err := manager.repository.GetAllAppointmentsByProfessionalId(ctx, professionalId)
	if err != nil {
		manager.logger.Error("Error get all appointments by professional owner", "use_case_manager", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error get all appointment by professional id"), errors.WithError(err))
	}

	manager.logger.Info("Appointments retrieved successfully.", "use_case_manager")

	return appointments, nil
}

func (manager *AppointmentManager) GetAppointmentById(ctx context.Context, appointmentId string) (*domain.Appointment, *errors.Exception) {
	appointment, err := manager.repository.GetAppointmentById(ctx, appointmentId)
	if err != nil {
		manager.logger.Error("Error get appointment detail", "use_case_manager", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error get appointment by id"), errors.WithError(err))
	}

	manager.logger.Info("Appointment retrieved successfully.", "use_case_manager")

	return appointment, nil
}

func (manager *AppointmentManager) DeleteAppointment(ctx context.Context, appointmentId string) *errors.Exception {
	err := manager.repository.DeleteAppointment(ctx, appointmentId)
	if err != nil {
		manager.logger.Error("Error deleting appointment", "use_case_manager", "err", err.Error())

		return errors.Unexpected(errors.WithMessage("error get appointment by id"), errors.WithError(err))
	}

	manager.logger.Info("Appointment deleted successfully.", "use_case_manager")

	return nil
}
