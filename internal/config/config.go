package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey string
	Model string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
		Model: os.Getenv("OPENAI_MODEL"),
	}
}