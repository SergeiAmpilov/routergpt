package rest

import (
	"github.com/gofiber/fiber/v2"
	"routergpt/internal/domain/completion/handler"
)

func SetupRoutes(app *fiber.App, h handler.Handler) {
	// Health check route
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// Completion route
	app.Post("/v1/completions", h.CreateCompletion)
}