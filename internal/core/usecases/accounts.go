package usecases

import (
	"context"
	"regexp"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/eventconstants"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/errors"
	"github.com/sirupsen/logrus"
)

type AccountManager struct {
	repository outbound.AccountRepository
	messaging  outbound.Publisher
	hasher     outbound.Hasher
	logger     *logrus.Logger
}

func NewAccount(
	repository outbound.AccountRepository,
	publisher outbound.Publisher,
	hasher outbound.Hasher,
	logger *logrus.Logger,
) inbound.AccountUseCase {
	return &AccountManager{
		repository: repository,
		messaging:  publisher,
		hasher:     hasher,
		logger:     logger,
	}
}

func (manager *AccountManager) Add(ctx context.Context, payload *domain.Account) (*domain.Account, *errors.Exception) {
	isValidEmail := validateAccountEmail(payload.Email)
	if !isValidEmail {
		manager.logger.Error("Error validating email.", "account_use_case_manager")

		return nil, errors.BadRequest(errors.WithMessage("invalid email"))
	}

	account, _ := manager.repository.FindAccountByEmail(ctx, payload.Email)
	if account != nil {
		return nil, errors.Confilct(errors.WithMessage("account already exists in the database"))
	}

	hash, err := manager.hasher.Hash(payload.Password)
	if err != nil {
		manager.logger.Error("Error generating password hash.", "account_use_case_manager", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error generating password hash"))
	}
	payload.Password = hash

	account, err = manager.repository.Add(ctx, payload)
	if err != nil {
		manager.logger.Error("Error saving account to repository.", "account_use_case_manager", "err", err.Error())

		return nil, errors.Unexpected()
	}

	e := inbound.Event{
		Type: eventconstants.AccountCreatedEventType,
		Payload: inbound.AccountCreatedEvent{
			Email: account.Email,
		},
	}

	err = manager.messaging.Publish(ctx, e)
	if err != nil {
		manager.logger.Println("Some error has been occurred publishing the message in broker.", "account_use_case_manager")
		// CONSISTENCY ISSUE: at this point the appointment was already persisted in the database
		// but the event failed to be published to the broker, leaving the system in an inconsistent state.
		// The consumer will never process this appointment, and retrying the request will result in a CONFLICT error.
		// TODO: implement the Outbox Pattern — persist the event in the same transaction as the appointment,
		// and use a background worker (relay) to publish it to the broker, ensuring at-least-once delivery.
		return nil, errors.Unexpected(errors.WithMessage("cannot publish the message in broker"))
	}

	manager.logger.Println("Account create and message was published.", "account_use_case_manager")

	return account, nil
}

func (manager *AccountManager) FindAccountById(ctx context.Context, id string) (*domain.Account, *errors.Exception) {
	account, err := manager.repository.FindAccountById(ctx, id)
	if err != nil {
		manager.logger.Error("Error finding account by id.", "account_use_case_manager", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error finding account"), errors.WithError(err))
	}

	manager.logger.Info("Account found successfully.", "account_use_case_manager")

	return account, nil
}

func (manager *AccountManager) FindAllAccounts(ctx context.Context) ([]domain.Account, *errors.Exception) {
	account, err := manager.repository.FindAllAccounts(ctx)
	if err != nil {
		manager.logger.Error("Error finding all accounts.", "account_use_case_manager", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("some error has been ocurred"))
	}

	manager.logger.Info("All accounts retrieved successfully.", "account_use_case_manager")

	return account, nil
}

func (manager *AccountManager) FindAllEstablishmentsByAccountId(ctx context.Context, accountId string) ([]domain.Establishment, *errors.Exception) {
	establishments, err := manager.repository.FindAllEstablishmentsByAccountId(ctx, accountId)
	if err != nil {
		manager.logger.Error("Error finding establishments by account id.", "account_use_case_manager", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("some error has been ocurred"))
	}

	manager.logger.Info("Establishments retrieved successfully.", "account_use_case_manager")

	return establishments, nil
}

func validateAccountEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
