package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	appConfig *AppConfig
}

func New() Config {
	godotenv.Load()

	cfg := &config{}
	cfg.loadFromEnv()

	return cfg

}

func (c *config) loadFromEnv() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	c.appConfig = &AppConfig{
		Port: port,
	}
}

func (c *config) GetConfig() *AppConfig {
	return c.appConfig
}
