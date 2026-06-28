package handler

import "github.com/gofiber/fiber/v2"

type completionHandler struct{}

func New() Handler {
	return &completionHandler{}
}

func (h *completionHandler) CreateCompletion(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "OK"})
}