package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"igaopk.com/goPower/internal/domain/repository"
	"igaopk.com/goPower/internal/service"
	"igaopk.com/goPower/internal/usecase"
)

type SeasonController interface {
	GetSeasons(context *gin.Context)
	GetSeason(context *gin.Context)
}

type seasonController struct {
	fetchAllSeasonsUseCase usecase.FetchAllSeasonsUseCase
	fetchSeasonByIDUseCase usecase.FetchSeasonByIDUseCase
}

// NewSeasonControllerWithDeps inicializa o controller com todas as dependências
func NewSeasonController(dbConn *gorm.DB) SeasonController {
	repository := repository.NewSeasonGormRepository(dbConn)

	service := service.NewSeasonService(repository)

	fetchAllSeasonsUseCase := usecase.NewFetchAllSeasonsUseCase(service)
	fetchSeasonByIDUseCase := usecase.NewFetchSeasonByIDUseCase(service)

	controller := &seasonController{
		fetchAllSeasonsUseCase: fetchAllSeasonsUseCase,
		fetchSeasonByIDUseCase: fetchSeasonByIDUseCase,
	}

	return controller
}

func (c *seasonController) GetSeasons(conntext *gin.Context) {
	seasons := c.fetchAllSeasonsUseCase.Execute()
	conntext.JSON(200, seasons)
}

func (c *seasonController) GetSeason(context *gin.Context) {

}
