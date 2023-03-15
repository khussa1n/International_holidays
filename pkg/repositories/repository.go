package repositories

import (
	"github.com/jmoiron/sqlx"
	"telegram_bot_golang/pkg/models"
)

type Users interface {
	CreateUser(user *models.Users) error
	UpdateUserAllQueriesCount(charID int64) error
	GetUserFirstQueryTime(userID int64) (string, error)
	GetUserAllQueriesCount(userID int64) (int, error)
	GetUserID(chatID int64) (int, error)
	CheckUserByID(userID int64) string
}

type Dates interface {
	CreateDate(date *models.Dates) error
	DeleteDate(userID string, description string) error
	GetDateByDate(date string) ([]string, error)
	GetAllDate(userID int64) ([]models.Dates, error)
}

type Repository struct {
	Users
	Dates
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUserRepo(db),
		Dates: NewDateRepo(db),
	}
}
