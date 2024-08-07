package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	SecretKey  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Config successfully loaded!")

	SecretKey = os.Getenv("SECRET_KEY")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
}
