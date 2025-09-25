package service

import (
	"igaopk.com/goPower/internal/models"
	"igaopk.com/goPower/internal/repository"
)

type RangerService struct {
	repo repository.RangerRepository
}

func NewRangerService(repo repository.RangerRepository) *RangerService {
	return &RangerService{repo: repo}
}

func (service *RangerService) GetAllRangers() ([]models.Ranger, error) {
	return service.repo.FindAll()
}
