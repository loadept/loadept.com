package config

type config struct {
	PORT      string `env:"PORT" default:"8080"`
	DB_NAME   string `env:"DB_NAME" default:"db.sqlite"`
	REDIS_URI string `env:"REDIS_URI" default:""`
}
