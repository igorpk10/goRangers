package usecase

import (
	"igaopk.com/goPower/internal/domain/models"
	"igaopk.com/goPower/internal/service"
)

type FetchSeasonByIDUseCase interface {
	Execute(id int) *models.Season
}

type fetchSeasonUseCase struct {
	service service.SeasonService
}

func NewFetchSeasonByIDUseCase(service service.SeasonService) FetchSeasonByIDUseCase {
	return &fetchSeasonUseCase{service: service}
}

func (uc fetchSeasonUseCase) Execute(id int) *models.Season {
	result, err := uc.service.GetSeasonByID(id)

	if err != nil {
		return nil
	}

	return result
}
