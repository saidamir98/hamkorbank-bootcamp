package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // development, staging, release

	LogLevel string
	HTTPPort string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7070"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "test"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "test"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
