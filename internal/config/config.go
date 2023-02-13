package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"hsr/loto/internal/logger"
	"os"
	"strconv"
)

const (
	defaultHttpPort = ":80"
	defaultGrpcPort = ":81"
)

// Config represents an application configuration.
type Config struct {
	HttpPort           string
	GrpcPort           string
	DSN                string
	ChatId             int
	SwaggerPath        string
	ExternalServiceUrl string
}

func Get() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Error.Printf("failed to load env vars: %s", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s database=%s user=%s password=%s",
		getEnv("DB_HOST", ""),
		getEnv("DB_PORT", ""),
		getEnv("DB_DATABASE", ""),
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
	)

	return &Config{
		HttpPort:           getEnv("HTTP_PORT", defaultHttpPort),
		GrpcPort:           getEnv("GRPC_PORT", defaultGrpcPort),
		DSN:                dsn,
		ChatId:             getEnvAsInt("CHAT_ID", 0),
		SwaggerPath:        getEnv("SWAGGER_PATH", "api/api.swagger.json"),
		ExternalServiceUrl: getEnv("EXTERNAL_SERVICE_URL", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if v, e := strconv.Atoi(value); e == nil {
		return v
	}

	return defaultValue
}
