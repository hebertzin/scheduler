package usecases

import (
	"context"
	"regexp"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/emailtemplates"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AccountManager struct {
	repository    domain.AccountRepository
	logger        *logrus.Logger
	emailProvider domain.EmailSender
}

func NewAccount(repository domain.AccountRepository, emailProvider domain.EmailSender, logger *logrus.Logger) domain.AccountUseCase {
	return &AccountManager{
		repository:    repository,
		emailProvider: emailProvider,
		logger:        logger,
	}
}

func (a *AccountManager) Add(ctx context.Context, payload *domain.Account) (*domain.Account, *core.Exception) {
	isValidEmail := validateAccountEmail(payload.Email)
	if !isValidEmail {
		return nil, core.BadRequest(core.WithMessage("invalid email"))
	}

	account, _ := a.repository.FindAccountByEmail(ctx, payload.Email)
	if account == nil {
		return nil, core.Confilct(core.WithMessage("account already exists in the database"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error generating password hash"))
	}
	payload.Password = string(hash)

	account, err = a.repository.Add(ctx, payload)
	if err != nil {
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

	a.emailProvider.Send(message)
	return account, nil
}

func (a *AccountManager) FindAccountById(ctx context.Context, id string) (*domain.Account, *core.Exception) {
	account, err := a.repository.FindAccountById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding account"), core.WithError(err))
	}

	return account, nil
}

func (a *AccountManager) FindAllAccounts(ctx context.Context) ([]domain.Account, *core.Exception) {
	account, err := a.repository.FindAllAccounts(ctx)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been ocurred"))
	}

	return account, nil
}

func (a *AccountManager) FindAllEstablishmentsByAccountId(ctx context.Context, accountId string) ([]domain.Establishment, *core.Exception) {
	establishments, err := a.repository.FindAllEstablishmentsByAccountId(ctx, accountId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been ocurred"))
	}

	return establishments, nil
}

func validateAccountEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
