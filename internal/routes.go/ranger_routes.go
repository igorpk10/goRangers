package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRangerRoutes(r *gin.Engine) {
	r.GET("/rangers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Lista de produtos"})
	})
}
