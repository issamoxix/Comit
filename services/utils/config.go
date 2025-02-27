package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAiKey string
	ApiUrl    string
}

var Version = "0.5"
var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		OpenAiKey: getEnv("OPENAI_KEY", "XXX-XXXX-XXX"),
		ApiUrl:    getEnv("API_URL", "http://localhost:8000/"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
