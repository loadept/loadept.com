package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string
	DB_NAME   string
	REDIS_URI string
}

var Env *Config

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

// TODO: apply singleton
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env variables defined, using default variables")
	}

	Env = &Config{
		PORT:      getEnv("PORT", "8080"),
		DB_NAME:   getEnv("DB_NAME", "database.db"),
		REDIS_URI: getEnv("REDIS_URI", ""),
	}
}
