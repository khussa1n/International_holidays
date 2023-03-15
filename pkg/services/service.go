package services

import (
	"telegram_bot_golang/pkg/models"
	"telegram_bot_golang/pkg/repositories"
)

type Users interface {
	CreateUser(user *models.Users) error
	UpdateUserAllQueriesCount(charID int64) error
	GetUserFirstQueryTime(userID int64) (string, error)
	GetUserAllQueriesCount(chatID int64) (int, error)
	GetUserID(chatID int64) (int, error)
	CheckUserByID(userID int64) string
}

type Dates interface {
	CreateDate(date *models.Dates) error
	DeleteDate(userID string, description string) error
	GetDateByDate(date string) ([]string, error)
	GetAllDate(userID int64) ([]models.Dates, error)
}

type Service struct {
	Users
	Dates
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Users: NewUserService(repository.Users),
		Dates: NewDateService(repository.Dates),
	}
}
