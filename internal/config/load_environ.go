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

// LoadConfig loads environment variables from a .env file, if present,
//
// or uses default values defined in the "config" structure.
//
// The function is only executed once thanks to "sync.Once", ensuring that
//
// the configuration is loaded safely and efficiently.
//
// Values are obtained using "env" and "default" tags in the "config" structure:
//
// - "env": Name of the environment variable.
//
// - "default": Default value if the environment variable is not defined.
//
// Example structure:
//
//	type config struct {
//		PORT string `env:"PORT" default:"8080"`
//		DB string `env:"DB_NAME" default:"database.db"`
//	}
//
// Once the configuration is loaded, the variables can be accessed via "Env".
func LoadEnviron() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env variables defined, using default variables")
		}

		Env = &config{}

		v := reflect.ValueOf(Env).Elem()
		t := v.Type()

		for i := range v.NumField() {
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
