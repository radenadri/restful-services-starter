package utils

import (
	"github.com/joho/godotenv"
	"os"
)

var env = godotenv.Load()

func LoadEnv(key string) string {
	if env != nil {
		panic("Failed to load .env file")
	}

	return os.Getenv(key)
}
