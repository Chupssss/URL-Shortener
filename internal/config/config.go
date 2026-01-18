package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("config: Файл .env не найден, используются переменные окружения")
	}

	return &Config{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaulVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaulVal
}
