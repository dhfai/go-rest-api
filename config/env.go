package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

var Envs = initConfig()

func initConfig() *Config {
	godotenv.Load()

	return &Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "your_db"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
		TimeZone: getEnv("DB_TIMEZONE", "Asia/Makassar"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
