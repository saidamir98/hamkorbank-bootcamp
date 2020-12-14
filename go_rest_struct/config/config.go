package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string // development, staging, production

	LogLevel string
	HttpPort string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7070"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
