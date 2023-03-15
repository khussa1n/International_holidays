package services

import (
	"telegram_bot_golang/pkg/models"
	"telegram_bot_golang/pkg/repositories"
)

type DateService struct {
	repo repositories.Dates
}

func NewDateService(repo repositories.Dates) *DateService {
	return &DateService{repo: repo}
}

func (s *DateService) CreateDate(date *models.Dates) error {
	return s.repo.CreateDate(date)
}

func (s *DateService) DeleteDate(userID string, description string) error {
	return s.repo.DeleteDate(userID, description)
}

func (s *DateService) GetDateByDate(date string) ([]string, error) {
	return s.repo.GetDateByDate(date)
}

func (s *DateService) GetAllDate(userID int64) ([]models.Dates, error) {
	return s.repo.GetAllDate(userID)
}
