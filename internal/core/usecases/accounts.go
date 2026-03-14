package usecases

import (
	"context"
	"regexp"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/smtp"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AccountUseCase struct {
	repository domain.AccountRepository
	logger     *logrus.Logger
	smptConfig *smtp.SMPTConfig
}

func NewAccountUseCase(
	repository domain.AccountRepository,
	logger *logrus.Logger,
	smtpConfig *smtp.SMPTConfig,
) domain.AccountUseCase {

	s := smtp.NewSMPT(smtpConfig.Port, smtpConfig.Password, smtpConfig.Host)

	return &AccountUseCase{
		repository: repository,
		logger:     logger,
		smptConfig: s,
	}
}

func (s *AccountUseCase) Add(ctx context.Context, payload *domain.Account) (*domain.Account, *core.Exception) {
	isValidEmail := validateAccountEmail(payload.Email)
	if !isValidEmail {
		return nil, core.BadRequest(core.WithMessage("invalid email"))
	}

	account, _ := s.repository.FindAccountByEmail(ctx, payload.Email)
	if account == nil {
		return nil, core.Confilct(core.WithMessage("account already exists in the database"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error generating password hash"))
	}
	payload.Password = string(hash)

	a, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	message := smtp.SMPTSendEmail{
		From:    "hebertsantosdeveloper@gmail.com",
		To:      []string{account.Email},
		Subject: "Confirm your account",
		Message: `Hello, Thank you for creating your account.
                  To activate your account, please confirm your email address by clicking the link below.
                  If you did not create this account, you can safely ignore this email.
                  Best regards.`,
	}

	// if smpt fails, dont break the user flow
	_ = s.smptConfig.Send(message)

	return a, nil
}

func (s *AccountUseCase) FindAccountById(ctx context.Context, accountId string) (*domain.Account, *core.Exception) {
	account, err := s.repository.FindAccountById(ctx, accountId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding account"), core.WithError(err))
	}

	return account, nil
}

func (s *AccountUseCase) FindAllAccounts(ctx context.Context) ([]domain.Account, *core.Exception) {
	account, err := s.repository.FindAllAccounts(ctx)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been ocurred"))
	}

	return account, nil
}

func (s *AccountUseCase) FindAllEstablishmentsByAccountId(ctx context.Context, accountId string) ([]domain.Establishment, *core.Exception) {
	establishments, err := s.repository.FindAllEstablishmentsByAccountId(ctx, accountId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been ocurred"))
	}

	return establishments, nil
}

func validateAccountEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
