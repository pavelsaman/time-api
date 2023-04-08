package config

import "os"

var defaultVersion string = "v1"
var defaultPort string = ":8080"

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
