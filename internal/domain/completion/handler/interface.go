package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateCompletion(c *fiber.Ctx) error
}