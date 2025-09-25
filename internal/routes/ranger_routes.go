package routes

import (
	"github.com/gin-gonic/gin"
	"igaopk.com/goPower/internal/controller"
)

func setupRangerRoutes(r *gin.Engine) {
	var rangerController = controller.InitRangerController()

	r.GET("/rangers", rangerController.GetAllRangers)
}
