package config

import (
	"os"
)

type Config struct {
	Port              string
	Environment       string
	ReadTimeout       int
	WriteTimeout      int
	JWTSecret         string
	JWTExpiration     int
	Database          DatabaseConfig
	Redis             RedisConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func Load() *Config {
	return &Config{
		Port:          getEnv("PORT", "8081"),
		Environment:   getEnv("ENVIRONMENT", "development"),
		ReadTimeout:   15,
		WriteTimeout:  15,
		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiration: 3600, // 1 hour
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "auth_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			TimeZone: getEnv("DB_TIMEZONE", "UTC"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
