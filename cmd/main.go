package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"igaopk.com/goPower/internal/config"
)

func main() {
	config.LoadEnvs()

	fmt.Println("Its works")
	router := gin.Default()

	port := config.GetEnv("APP_PORT", "8080")

	router.GET("/health", func(ctx *gin.Context) {
		user := config.GetEnv("DB_USER", "Error")
		ctx.JSON(200, gin.H{"status": "We are the power " + user})
	})

	router.Run(":" + port)
}
