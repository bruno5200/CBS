package handler

import (
	p "github.com/bruno5200/CSM/service/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
)

func (h *serviceHandler) GetAll(c *fiber.Ctx) error {

	services, err := h.serviceService.GetServices()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.ServiceErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.ServicesSuccessResponse(services))
}
