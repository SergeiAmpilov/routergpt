package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	appConfig *AppConfig
}

func New() Config {
	// Load .env file if it exists
	godotenv.Load()
	
	cfg := &config{}
	cfg.loadFromEnv()
	return cfg
}

func (c *config) loadFromEnv() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}
	
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}
	
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = ""
	}
	
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}
	
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)
	
	c.appConfig = &AppConfig{
		Port:       port,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBSSLMode:  dbSSLMode,
		ConnString: connString,
	}
}

func (c *config) GetConfig() *AppConfig {
	return c.appConfig
}