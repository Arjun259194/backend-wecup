package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error while loading environment variables - %v", err)
	}
	return os.Getenv(key)
}

var (
	PORT              string = getEnv("PORT")
	DB_CONNECTION_URI string = getEnv("DB_CONNECTION_URI")
	SECRET_TOKEN      string = getEnv("SECRET_TOKEN")
)
