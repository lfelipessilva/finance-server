package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost              string `env:"DB_HOST" envDefault:"localhost"`
	DBPort              string `env:"DB_PORT" envDefault:"5432"`
	DBUser              string `env:"DB_USER" envDefault:"postgres"`
	DBPassword          string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName              string `env:"DB_NAME" envDefault:"finance"`
	SSLMode             string `env:"SSL_MODE" envDefault:"disable"`
	JWTSecret           string `env:"JWT_SECRET" envDefault:"your-secret-key"`
	GoogleOAuthClientID string `env:"GOOGLE_OAUTH_CLIENT_ID" envDefault:""`
}

func Load() (*Config, error) {
	// Try to load .env file from the project root
	projectRoot, err := findProjectRoot()
	if err != nil {
		log.Printf("Warning: Could not find project root: %v", err)
	} else {
		envPath := filepath.Join(projectRoot, ".env")
		if err := godotenv.Load(envPath); err != nil {
			log.Printf("Warning: Could not load .env file from %s: %v", envPath, err)
		} else {
			log.Printf("Successfully loaded .env file from %s", envPath)
		}
	}

	// Also try loading from current directory as fallback
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file from current directory: %v", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	// Log the loaded configuration (without sensitive data)
	log.Printf("Loaded config - DB_HOST: %s, DB_PORT: %s, DB_NAME: %s, SSL_MODE: %s",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.SSLMode)

	return cfg, nil
}

// findProjectRoot looks for the project root by finding the go.mod file
func findProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree looking for go.mod
	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			return "", fmt.Errorf("could not find go.mod file")
		}
		currentDir = parent
	}
}
