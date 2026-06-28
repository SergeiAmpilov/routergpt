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

	return c.JSON(req)
}