package main

import (
	"github.com/gin-gonic/gin"
	"igaopk.com/goPower/internal/config"
	"igaopk.com/goPower/internal/database"
	"igaopk.com/goPower/internal/logger"
	"igaopk.com/goPower/internal/routes.go"
)

func main() {
	config.LoadEnvs()
	logger.LogMessage("Starting the application...", logger.INFO)
	logger.LogMessage("Environment variables loaded.", logger.INFO)

	var db = config.GetEnv("DB_NAME")
	logger.LogMessage("Connecting to the database %s ...", logger.INFO, db)

	database.Connect()
	logger.LogMessage("Database connected successfully.", logger.INFO)

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":" + config.GetEnv("APP_PORT"))
}
