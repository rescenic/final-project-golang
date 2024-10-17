// config/config.go

package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var DB *sql.DB

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

	// Initialize database connection
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	fmt.Println("Database connected successfully")

	return config
}

func getEnvOrFail(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Error: %s environment variable not set", key)
	}
	return value
}
