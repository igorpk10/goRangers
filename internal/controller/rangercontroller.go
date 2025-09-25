package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"igaopk.com/goPower/internal/database"
	"igaopk.com/goPower/internal/repository"
	"igaopk.com/goPower/internal/service"
	"igaopk.com/goPower/internal/usecase"
)

type RangerController struct {
	getAllRangersUseCase *usecase.GetAllRangersUseCase
}

func InitRangerController() *RangerController {
	repo := repository.NewRangerGormRepository(database.DB)
	svc := service.NewRangerService(repo)
	uc := usecase.NewGetAllRangersUseCase(svc)
	return newRangerController(uc)
}

func newRangerController(getAllUC *usecase.GetAllRangersUseCase) *RangerController {
	return &RangerController{
		getAllRangersUseCase: getAllUC,
	}
}

func (ctrl *RangerController) GetAllRangers(c *gin.Context) {
	rangers, err := ctrl.getAllRangersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rangers)
}
