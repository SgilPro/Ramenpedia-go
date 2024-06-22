package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

var Environment string

func LoadEnvVariables() {

	err := godotenv.Load(".env." + Environment)
	if err != nil {
		log.Fatal("Error loading .env file, Environment=" + Environment)
	}
}
