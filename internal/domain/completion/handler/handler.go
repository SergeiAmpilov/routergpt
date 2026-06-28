package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"routergpt/internal/domain/completion/model"
)

type completionHandler struct{}

func New() Handler {
	return &completionHandler{}
}

func (h *completionHandler) CreateCompletion(c *fiber.Ctx) error {
	var req model.CompletionRequestDTO
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(req)
}