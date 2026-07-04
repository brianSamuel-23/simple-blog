package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Failed to load .env file: %v. Falling back to OS environment variables.\n", err)
	}
}

func getEnv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("Warning: Environment variable %s is not set, using default value %q.\n", key, defaultValue)
		return defaultValue
	}

	return val
}

func getEnvBool(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	parsed, err := strconv.ParseBool(val)
	if err != nil {
		log.Printf("Warning: Environment variable %s is not a valid boolean, using default value %t.\n", key, defaultValue)
		return defaultValue
	}

	return parsed
}

func getEnvInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("Warning: Environment variable %s is not a valid integer, using default value %d.\n", key, defaultValue)
		return defaultValue
	}

	return parsed
}
