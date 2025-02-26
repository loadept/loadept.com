package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type config struct {
	PORT      string
	DB_NAME   string
	REDIS_URI string
}

var (
	Env  *config
	once sync.Once
)

func LoadConfig() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env variables defined, using default variables")
		}

		Env = &config{
			PORT:      getEnv("PORT", "8080"),
			DB_NAME:   getEnv("DB_NAME", "database.db"),
			REDIS_URI: getEnv("REDIS_URI", ""),
		}
	})
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
