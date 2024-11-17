package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Not loaded env vars")
	}
}
