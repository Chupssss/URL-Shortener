package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	BaseURL string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("config: Файл .env не найден, используются переменные окружения")
	}

	return &Config{
		Port:    getEnv("PORT", "8080"),
		BaseURL: getEnv("BASE_URL", "http://localhost:8080/"),
		DBHost:  getEnv("DB_HOST", "loacalhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "youruser"),
		DBPass:  getEnv("DB_PASSWORD", "yourpassword"),
		DBName:  getEnv("DB_NAME", "yourdbname"),
	}
}

func getEnv(key, defaulVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaulVal
}
