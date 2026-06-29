package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateModel(c *fiber.Ctx) error
}