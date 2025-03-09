package config

// config defines the structure for storing the application's configuration variables.
//
// Each field represents an environment variable and may include tags for configuration:
// - "env": Specifies the name of the environment variable to load.
// - "default": Defines a default value if the environment variable is not set.
//
// Usage example:
//
//	type config struct {
//		PORT      string `env:"PORT" default:"8080"`
//		DB_NAME   string `env:"DB_NAME" default:"database.db"`
//		REDIS_URI string `env:"REDIS_URI" default:""`
//	}
//
// Once the configuration is loaded using "LoadConfig()", values can be accessed via "Env":
//
//	fmt.Println(Env.PORT)  // Prints the configured port or "8080" if not set.
type config struct {
	PORT      string `env:"PORT" default:"8080"`
	DB_NAME   string `env:"DB_NAME" default:"db.sqlite"`
	REDIS_URI string `env:"REDIS_URI" default:""`
	SECRET_KEY string `env:"SECRET_KEY" default:""`
}
