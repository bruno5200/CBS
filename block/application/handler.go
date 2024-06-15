package application

import "github.com/gofiber/fiber/v2"

type BlockHandler interface {
	Get(c *fiber.Ctx) error
	Post(c *fiber.Ctx) error
	Json(c *fiber.Ctx) error
	Put(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetByGroup(c *fiber.Ctx) error
	GetByService(c *fiber.Ctx) error
}
