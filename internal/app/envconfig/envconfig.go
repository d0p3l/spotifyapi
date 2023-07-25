package envconfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SPOTIFY_ID     string
	SPOTIFY_SECRET string
}

func New() *Config {
	return &Config{
		SPOTIFY_ID:     getEnv("SPOTIFY_ID", ""),
		SPOTIFY_SECRET: getEnv("SPOTIFY_SECRET", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	err := godotenv.Load("/.env")
	if err != nil {
		log.Print("No .env file found")
	}

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
