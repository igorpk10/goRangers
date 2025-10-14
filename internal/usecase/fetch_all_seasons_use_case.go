package usecase

import (
	"net/http"

	"igaopk.com/goPower/internal/domain/models"
	"igaopk.com/goPower/internal/service"
)

type FetchAllSeasonsUseCase interface {
	Execute() *[]models.Season
}

type fetchAllSeasonsUseCase struct {
	service service.SeasonService
}

func NewFetchAllSeasonsUseCase(service service.SeasonService) FetchAllSeasonsUseCase {
	return &fetchAllSeasonsUseCase{service: service}
}

func (uc *fetchAllSeasonsUseCase) Execute() *[]models.Season {
	result, err := uc.service.GetAllSeasons()

	if len(*result) == 0 {
		panic(models.HTTPError{
			StatusCode: http.StatusNoContent,
			Message:    "No seasons found",
			Error:      err.Error(),
		})
	}

	return result
}
