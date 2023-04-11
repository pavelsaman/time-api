package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var defaultVersion string = "v1"
var defaultPort string = ":8080"

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading .env file")
	}
}

func ApiVersion() string {
	if apiVersion := os.Getenv("API_VERSION"); apiVersion != "" {
		return apiVersion
	}

	return defaultVersion
}

func ApiPort() string {
	if apiPort := os.Getenv("API_PORT"); apiPort != "" {
		return ":" + apiPort
	}

	return defaultPort
}
