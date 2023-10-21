package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
