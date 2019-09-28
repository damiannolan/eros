package server

import (
	"os"
	"strconv"
)

// Environment Vars
const (
	EnvMetricsEnabled = "METRICS_ENABLED"
	EnvServerHost     = "SERVER_HOST"
	EnvServerPort     = "SERVER_PORT"
)

// Config exposes Server configuration options
type Config struct {
	BasePath       string
	MetricsEnabled bool
	Name           string
	Port           int
}

// NewConfig constructs a new *server.Config instance with defaults
func NewConfig() *Config {
	return &Config{
		BasePath:       "/api/v1",
		Name:           getStringEnvironmentVariable(EnvServerHost, "eros-http"),
		Port:           getIntEnvironmentVariable(EnvServerPort, 8080),
		MetricsEnabled: getBoolEnvironmentVariable(EnvMetricsEnabled, false),
	}
}

func getBoolEnvironmentVariable(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		boolValue, _ := strconv.ParseBool(value)
		return boolValue
	}
	return fallback
}

func getIntEnvironmentVariable(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		intValue, _ := strconv.Atoi(value)
		return intValue
	}
	return fallback
}

func getStringEnvironmentVariable(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
