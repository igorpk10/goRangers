package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	err := godotenv.Load("../.env") // sobe uma pasta e busca o .env
	if err != nil {
		log.Println(".env file not found or failed to load (defaulting to system env)")
	}
}

func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	panic("Environment variable " + key + " not set")
}
