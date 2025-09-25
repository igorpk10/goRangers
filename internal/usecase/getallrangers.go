package usecase

import (
	"igaopk.com/goPower/internal/models"
	"igaopk.com/goPower/internal/service"
)

type GetAllRangersUseCase struct {
	rangerService *service.RangerService
}

func NewGetAllRangersUseCase(rangerService *service.RangerService) *GetAllRangersUseCase {
	return &GetAllRangersUseCase{rangerService: rangerService}
}

func (uc *GetAllRangersUseCase) Execute() ([]models.Ranger, error) {
	return uc.rangerService.GetAllRangers()
}
