package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAiKey string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		OpenAiKey: getEnv("OPENAI_KEY", "XXX-XXXX-XXX"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
