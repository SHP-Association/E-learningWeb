package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
	JWT        JWTConfig
	CORS       CORSConfig
	Upload     UploadConfig
	Email      EmailConfig
	Pagination PaginationConfig
	RateLimit  RateLimitConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host            string
	Port            string
	Name            string
	User            string
	Password        string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type JWTConfig struct {
	Secret     string
	Expiry     time.Duration
	BcryptCost int
}

type CORSConfig struct {
	AllowedOrigins []string
}

type UploadConfig struct {
	MaxSize    int64
	MediaPath  string
	StaticPath string
}

type EmailConfig struct {
	Host     string
	Port     int
	UseTLS   bool
	User     string
	Password string
	From     string
}

type PaginationConfig struct {
	DefaultPageSize int
	MaxPageSize     int
}

type RateLimitConfig struct {
	AnonLimit int
	UserLimit int
	Window    time.Duration
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("GO_PORT", "8002"),
			Env:  getEnv("GO_ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:            getEnv("GO_DB_HOST", "db"),
			Port:            getEnv("GO_DB_PORT", "5433"),
			Name:            getEnv("GO_DB_NAME", "shplearner"),
			User:            getEnv("GO_DB_USER", "sandesh"),
			Password:        getEnv("GO_DB_PASSWORD", "password123"),
			SSLMode:         getEnv("GO_DB_SSLMODE", "disable"),
			MaxOpenConns:    getEnvAsInt("GO_DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("GO_DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvAsDuration("GO_DB_CONN_MAX_LIFETIME", "5m"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("GO_JWT_SECRET", "go_pagoda_jwt_secret_key_change_in_production"),
			Expiry:     getEnvAsDuration("GO_JWT_EXPIRY", "24h"),
			BcryptCost: getEnvAsInt("GO_BCRYPT_COST", 10),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvAsSlice("GO_CORS_ALLOWED_ORIGINS", []string{"http://localhost:5173"}),
		},
		Upload: UploadConfig{
			MaxSize:    getEnvAsInt64("GO_UPLOAD_MAX_SIZE", 10485760), // 10MB
			MediaPath:  getEnv("GO_MEDIA_PATH", "/data/media"),
			StaticPath: getEnv("GO_STATIC_PATH", "/data/static"),
		},
		Email: EmailConfig{
			Host:     getEnv("GO_EMAIL_HOST", "smtp.example.com"),
			Port:     getEnvAsInt("GO_EMAIL_PORT", 587),
			UseTLS:   getEnvAsBool("GO_EMAIL_USE_TLS", true),
			User:     getEnv("GO_EMAIL_USER", "example@example.com"),
			Password: getEnv("GO_EMAIL_PASSWORD", "examplepassword"),
			From:     getEnv("GO_EMAIL_FROM", "example@example.com"),
		},
		Pagination: PaginationConfig{
			DefaultPageSize: getEnvAsInt("GO_DEFAULT_PAGE_SIZE", 20),
			MaxPageSize:     getEnvAsInt("GO_MAX_PAGE_SIZE", 100),
		},
		RateLimit: RateLimitConfig{
			AnonLimit: getEnvAsInt("GO_RATE_LIMIT_ANON", 100),
			UserLimit: getEnvAsInt("GO_RATE_LIMIT_USER", 1000),
			Window:    getEnvAsDuration("GO_RATE_LIMIT_WINDOW", "1h"),
		},
	}

	return cfg, nil
}

// GetDSN returns the database connection string
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode,
	)
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue string) time.Duration {
	valueStr := getEnv(key, defaultValue)
	if duration, err := time.ParseDuration(valueStr); err == nil {
		return duration
	}
	// Try to parse as default
	if duration, err := time.ParseDuration(defaultValue); err == nil {
		return duration
	}
	return 0
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	return strings.Split(valueStr, ",")
}
