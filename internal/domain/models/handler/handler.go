package handler

import (
	"routergpt/internal/domain/models/model"
	"routergpt/internal/domain/models/service"

	"github.com/gofiber/fiber/v2"
)

type modelsHandler struct {
	service service.ModelsService
}

func New(service service.ModelsService) Handler {
	return &modelsHandler{
		service: service,
	}
}

func (h *modelsHandler) CreateModel(c *fiber.Ctx) error {
	req := new(model.CreateAIModelRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
			"details": err.Error(),
		})
	}

	result, err := h.service.CreateModel(c.Context(), *req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}