package config

import (
	"os"
	"time"
)

type Config struct {
	Environment string
	Port        string

	// Database configuration
	Database DatabaseConfig

	// Redis configuration
	Redis RedisConfig

	// JWT configuration
	JWT JWTConfig

	// GAuth Core configuration
	GAuthCore GAuthCoreConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret     string
	Expiry     time.Duration
	RefreshExp time.Duration
}

type GAuthCoreConfig struct {
	ServerURL string
}

func Load() *Config {
	return &Config{
		Environment: getEnv("GIN_MODE", "development"),
		Port:        getEnv("PORT", "8080"),

		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "gauth"),
			Password: getEnv("DB_PASSWORD", "gauth_password"),
			Name:     getEnv("DB_NAME", "gauth_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},

		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		},

		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your_jwt_secret_change_in_production"),
			Expiry:     24 * time.Hour,
			RefreshExp: 7 * 24 * time.Hour,
		},

		GAuthCore: GAuthCoreConfig{
			ServerURL: getEnv("GAUTH_SERVER_URL", "http://localhost:9090"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
