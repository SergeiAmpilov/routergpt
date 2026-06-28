package main

import (
	"log"
	"routergpt/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.New().GetConfig()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! This is a Fiber web server.")
	})

	log.Fatal(app.Listen(":" + cfg.Port))
}
