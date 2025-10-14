package routes

import (
	"github.com/gin-gonic/gin"
	"igaopk.com/goPower/internal/controller"
	"igaopk.com/goPower/internal/database"
)

func setupRangerRoutes(r *gin.Engine) {

	var seasonController = controller.NewSeasonController(database.DB)
	r.GET("/seasons", seasonController.GetSeasons)
	r.GET("/seasons/:id", seasonController.GetSeason)

}
