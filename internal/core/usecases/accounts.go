package usecases

import (
	"context"
	"regexp"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/infra/emailtemplates"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AccountManager struct {
	repository    outbound.AccountRepository
	logger        *logrus.Logger
	emailProvider outbound.EmailSender
}

func NewAccount(
	repository outbound.AccountRepository,
	emailProvider outbound.EmailSender,
	logger *logrus.Logger,
) inbound.AccountUseCase {
	return &AccountManager{
		repository:    repository,
		emailProvider: emailProvider,
		logger:        logger,
	}
}

func (manager *AccountManager) Add(ctx context.Context, payload *domain.Account) (*domain.Account, *core.Exception) {
	isValidEmail := validateAccountEmail(payload.Email)
	if !isValidEmail {
		manager.logger.Error("Error validating email.", "account_use_case_manager")

		return nil, core.BadRequest(core.WithMessage("invalid email"))
	}

	account, _ := manager.repository.FindAccountByEmail(ctx, payload.Email)
	if account != nil {
		return nil, core.Confilct(core.WithMessage("account already exists in the database"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		manager.logger.Error("Error generating password hash.", "account_use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error generating password hash"))
	}
	payload.Password = string(hash)

	account, err = manager.repository.Add(ctx, payload)
	if err != nil {
		manager.logger.Error("Error saving account to repository.", "account_use_case_manager", "err", err.Error())

		return nil, core.Unexpected()
	}

	accountCreatedData := emailtemplates.AccountCreatedData{
		Email: account.Email,
	}

	body, _ := emailtemplates.RenderAccountCreated(accountCreatedData)

	message := domain.EmailMessage{
		From:    "hebertsantosdeveloper@gmail.com",
		To:      []string{account.Email},
		Subject: emailtemplates.AccountCreatedSubject,
		Message: body,
	}

	// this is sync (change to async later)
	manager.emailProvider.Send(message)

	manager.logger.Println("Account create and email sent successfully.", "account_use_case_manager")

	return account, nil
}

func (manager *AccountManager) FindAccountById(ctx context.Context, id string) (*domain.Account, *core.Exception) {
	account, err := manager.repository.FindAccountById(ctx, id)
	if err != nil {
		manager.logger.Error("Error finding account by id.", "account_use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error finding account"), core.WithError(err))
	}

	manager.logger.Info("Account found successfully.", "account_use_case_manager")

	return account, nil
}

func (manager *AccountManager) FindAllAccounts(ctx context.Context) ([]domain.Account, *core.Exception) {
	account, err := manager.repository.FindAllAccounts(ctx)
	if err != nil {
		manager.logger.Error("Error finding all accounts.", "account_use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("some error has been ocurred"))
	}

	manager.logger.Info("All accounts retrieved successfully.", "account_use_case_manager")

	return account, nil
}

func (manager *AccountManager) FindAllEstablishmentsByAccountId(ctx context.Context, accountId string) ([]domain.Establishment, *core.Exception) {
	establishments, err := manager.repository.FindAllEstablishmentsByAccountId(ctx, accountId)
	if err != nil {
		manager.logger.Error("Error finding establishments by account id.", "account_use_case_manager", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("some error has been ocurred"))
	}

	manager.logger.Info("Establishments retrieved successfully.", "account_use_case_manager")

	return establishments, nil
}

func validateAccountEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
