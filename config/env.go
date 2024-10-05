// Package config provides functionality to load and manage application configuration from environment variables.
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application's configuration values.
type Config struct {
	ServerHost string
	ServerPort string
}

// Envs holds the application's configuration loaded from environment variables.
var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, proceeding with defaults.")
	}

	return Config{
		ServerHost: getEnv("PUBLIC_HOST", "localhost"), // Default value to localhost
		ServerPort: getEnv("PORT", "8080"),             // Default value to 8080
	}
}

// getEnv retrieves the value of an environment variable or returns a fallback value if the variable is not set.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Printf("Warning: %s is not set, using fallback: %s", key, fallback)
	return fallback
}
