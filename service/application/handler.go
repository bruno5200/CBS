package application

import "github.com/gofiber/fiber/v2"

type ServiceHandler interface {
	Get(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Post(c *fiber.Ctx) error
	Put(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
