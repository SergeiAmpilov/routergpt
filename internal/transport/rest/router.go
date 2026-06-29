package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	completionhandler "routergpt/internal/domain/completion/handler"
	modelshandler "routergpt/internal/domain/models/handler"
)

func SetupRoutes(app *fiber.App, ch completionhandler.Handler, mh modelshandler.Handler) {
	// Health check route
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	// Completion route
	app.Post("/v1/completions", ch.CreateCompletion)

	// Models routes
	modelsGroup := app.Group("/models")
	modelsGroup.Post("/", mh.CreateModel)
}