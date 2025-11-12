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
	DEBUG           string `env:"DEBUG" default:"true"`
	PORT            string `env:"PORT" default:"8080"`
	DB_NAME         string `env:"DB_NAME" default:"db.sqlite"`
	REDIS_USER      string `env:"REDIS_USER" default:"default"`
	REDIS_HOST      string `env:"REDIS_HOST" default:"localhost"`
	REDIS_PORT      string `env:"REDIS_PORT" default:"6379"`
	REDIS_PASSWORD  string `env:"REDIS_PASSWORD" default:""`
	SECRET_KEY      string `env:"SECRET_KEY" default:""`
	GITHUB_API      string `env:"GITHUB_API" default:""`
	GITHUB_TOKEN    string `env:"GITHUB_TOKEN" default:""`
	PDF_SERVICE_URL string `env:"PDF_SERVICE_URL" default:""`
}
