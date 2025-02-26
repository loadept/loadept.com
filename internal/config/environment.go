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

// LoadConfig loads environment variables from a .env file
//
// if no such file is found, it will use default variables or
//
// system-defined variables.
//
// Prepares Env to load and directly access these variables
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

// getEnv loads system environment variables if not defined, loads default variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
