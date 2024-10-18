// config/config.go

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var DB *gorm.DB

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		DBHost:     getEnvOrFail("DB_HOST"),
		DBPort:     getEnvOrFail("DB_PORT"),
		DBUser:     getEnvOrFail("DB_USER"),
		DBPassword: getEnvOrFail("DB_PASSWORD"),
		DBName:     getEnvOrFail("DB_NAME"),
		JWTSecret:  getEnvOrFail("JWT_SECRET"),
	}

	// Initialize GORM database connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db
	fmt.Println("Database connected successfully using GORM")

	return config
}

func getEnvOrFail(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Error: %s environment variable not set", key)
	}
	return value
}
