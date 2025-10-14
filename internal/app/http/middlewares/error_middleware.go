package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"igaopk.com/goPower/internal/models"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				switch e := r.(type) {
				case models.HTTPError:
					// Se for 204, não manda nada no corpo
					if e.StatusCode == 204 {
						c.Status(204)
						return
					}

					// Para outros, manda JSON
					c.AbortWithStatusJSON(e.StatusCode, gin.H{
						"error": e.Message,
					})

				default:
					// Se não for nosso tipo, retorna 500 genérico
					log.Printf("Panic inesperado: %v", r)
					c.AbortWithStatusJSON(500, gin.H{
						"error": "Erro interno inesperado",
					})
				}
			}
		}()
		c.Next()
	}
}
