package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the application configuration
type ConfigPort struct {
	Port int
}

// NewConfig initializes a new Config instance with default values
func NewConfigPort() *ConfigPort {
	return &ConfigPort{
		Port: 8080, // Default port
	}
}

// LoadFromEnv loads configuration settings from environment variables
func (c *ConfigPort) LoadFromEnv() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Check if a port is set as an environment variable
	if portStr, ok := os.LookupEnv("APP_PORT"); ok {
		if port, err := strconv.Atoi(portStr); err == nil {
			c.Port = port
		}
	}

	return nil
}