package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envRegistry struct {
	DB_DSN     string
	JWT_Secret string
}

var envRegistryInstance *envRegistry

func loadDotenv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("[RestAPI]: Error in loading .env file: ", err)
	}

	log.Println("[RestAPI]: Successfully loaded info from .env")
}

func Config() *envRegistry {

	if envRegistryInstance == nil {

		loadDotenv()

		envRegistryInstance = &envRegistry{
			DB_DSN:     os.Getenv("DB_DSN"),
			JWT_Secret: os.Getenv("JWT_SECRET"),
		}
	}

	return envRegistryInstance
}
