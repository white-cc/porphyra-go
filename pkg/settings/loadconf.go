package settings

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConf() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load .env file")
	}
}
