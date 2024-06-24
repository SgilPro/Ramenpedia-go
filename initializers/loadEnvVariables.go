package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

var Environment string

func LoadEnvVariables() {
	log.Printf("[Environment=%v]\n", Environment)

	var err error
	if Environment != "" {
		err = godotenv.Load(".env." + Environment)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("Error loading .env file, Environment=" + Environment)
	}
}
