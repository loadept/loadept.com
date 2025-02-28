package config

import (
	"log"
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
)

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
// Prepares Env to load and directly access these variables.
func LoadConfig() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env variables defined, using default variables")
		}

		Env = &config{}

		v := reflect.ValueOf(Env).Elem()
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)

			envKey := field.Tag.Get("env")
			if envKey == "" {
				continue
			}

			defaultValue := field.Tag.Get("default")
			if defaultValue == "" {
				defaultValue = v.Field(i).String()
			}

			value := getEnv(envKey, defaultValue)
			v.Field(i).SetString(value)
		}
	})
}

// getEnv loads system environment variables if not defined, loads default variables.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
