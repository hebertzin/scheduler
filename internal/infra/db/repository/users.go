package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) *UserDatabaseRepository {
	return &UserDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *UserDatabaseRepository) Add(ctx context.Context, user *domains.User) (*domains.User, error) {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserDatabaseRepository) FindUserById(ctx context.Context, id string) (*domains.User, error) {
	var user *domains.User
	if err := repo.db.WithContext(ctx).
		Model(&models.Users{}).
		Where("id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserDatabaseRepository) FindUserByEmail(ctx context.Context, email string) (*domains.User, error) {
	var user *domains.User
	if err := repo.db.WithContext(ctx).
		Model(&models.Users{}).
		Where("email = ?", email).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserDatabaseRepository) FindAllUsers(ctx context.Context) ([]domains.User, error) {
	var users []domains.User
	err := repo.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserDatabaseRepository) FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]domains.Establishment, error) {
	var establishments []domains.Establishment
	err := repo.db.WithContext(ctx).
		Where("user_id = ?", user_id).
		Find(&establishments).Error

	if err != nil {
		return nil, err
	}
	return establishments, nil
}
