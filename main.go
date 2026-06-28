package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"routergpt/internal/config"
	"routergpt/internal/domain/completion/handler"
	"routergpt/internal/transport/rest"
)

func main() {

	cfg := config.New().GetConfig()

	app := fiber.New()

	h := handler.New()
	rest.SetupRoutes(app, h)

	log.Fatal(app.Listen(":" + cfg.Port))
}