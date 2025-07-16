package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	envFile := ".env." + env
	if _, err := os.Stat(envFile); err == nil {
		if err := godotenv.Load(envFile); err != nil {
			log.Printf("Warning: Could not load %s: %v", envFile, err)
		}
	} else {

		if err := godotenv.Load(); err != nil {
			log.Println("Warning: No .env file found, using environment variables")
		}
	}
}

func ValidateEnv() {
	required := []string{
		"APP_PORT",
		"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME",
		"REDIS_HOST", "REDIS_PORT",
		"API_KEY",
		"SECRET_KEY",
	}

	missing := []string{}
	for _, key := range required {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		log.Fatalf("Missing required environment variables: %v", missing)
	}

	log.Println("✅ Environment variables validated successfully!")
}

func InitializeAll() {
	log.Println("🚀 Initializing application...")
	LoadEnv()
	ValidateEnv()
	ConnectDatabase()
	ConnectRedis()
	RunMigrationsWithSeed()
	log.Println("🎉 Application initialized successfully!")
}
