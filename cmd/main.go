package main

import (
	"github.com/gin-gonic/gin"
	"igaopk.com/goPower/internal/config"
	"igaopk.com/goPower/internal/database"
	"igaopk.com/goPower/internal/logger"
)

func main() {
	logger.LogMessage("Starting the application...", logger.INFO)
	config.LoadEnvs()
	database.Connect()
	router := gin.Default()

	router.Run(":" + config.GetEnv("PORT", "8080"))
}
