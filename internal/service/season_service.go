package service

import (
	"igaopk.com/goPower/internal/domain/models"
	"igaopk.com/goPower/internal/domain/repository"
)

type SeasonService interface {
	GetAllSeasons() (*[]models.Season, error)
	GetSeasonByID(id int) (*models.Season, error)
}

type seasonService struct {
	repo repository.SeasonRepository
}

func NewSeasonService(repo repository.SeasonRepository) SeasonService {
	return &seasonService{repo: repo}
}

func (s *seasonService) GetAllSeasons() (*[]models.Season, error) {
	seasons, err := s.repo.FindAll()

	if len(*seasons) == 0 {
		return nil, 
	}

	return seasons, err
}

func (s *seasonService) GetSeasonByID(id int) (*models.Season, error) {
	season, err := s.repo.FindByID(id)
	return season, err
}
