package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig *Config

type Postgres struct {
	Host     string
	User     string
	Password string
	DBName   string
}

type Config struct {
	Postgres *Postgres
	Port     string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file: " + err.Error())
	}

	// Check for the presence of essential variables
	if err := checkEnvVariable("DB_HOST"); err != nil {
		return nil, err
	}
	if err := checkEnvVariable("DB_USER"); err != nil {
		return nil, err
	}
	if err := checkEnvVariable("DB_PASSWORD"); err != nil {
		return nil, err
	}
	if err := checkEnvVariable("DB_NAME"); err != nil {
		return nil, err
	}
	if err := checkEnvVariable("PORT"); err != nil {
		return nil, err
	}

	// Initialize your global variable with values from .env
	AppConfig = &Config{
		Postgres: &Postgres{
			Host:     getEnv("DB_HOST", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", ""),
		},
		Port: getEnv("PORT", "8080"),
	}

	// You can add more initialization logic as needed

	return AppConfig, nil
}

// Helper function to check if environment variable is present
func checkEnvVariable(key string) error {
	if _, exists := os.LookupEnv(key); !exists {
		return errors.New(key + " is not set in the .env file. Please provide a value.")
	}
	return nil
}

// Helper function to get environment variable or default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
