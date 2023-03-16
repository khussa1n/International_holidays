package services

import (
	"telegram_bot_golang/pkg/models"
	"telegram_bot_golang/pkg/repositories"
)

type UserService struct {
	repo repositories.Users
}

func NewUserService(repo repositories.Users) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.Users) error {
	_, err := s.repo.GetUserID(user.ChatID)
	if err != nil {
		return s.repo.CreateUser(user)
	}

	return nil
}

func (s *UserService) UpdateUserAllQueriesCount(charID int64) error {
	return s.repo.UpdateUserAllQueriesCount(charID)
}

func (s *UserService) GetUserFirstQueryTime(userID int64) (string, error) {
	return s.repo.GetUserFirstQueryTime(userID)
}

func (s *UserService) GetUserAllQueriesCount(chatID int64) (int, error) {
	return s.repo.GetUserAllQueriesCount(chatID)
}

func (s *UserService) GetUserID(chatID int64) (int, error) {
	return s.repo.GetUserID(chatID)
}

func (s *UserService) CheckUserByID(userID int64) string {
	return s.repo.CheckUserByID(userID)
}
