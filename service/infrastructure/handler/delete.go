package handler

import (
	d "github.com/bruno5200/CSM/service/domain"
	p "github.com/bruno5200/CSM/service/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *serviceHandler) Delete(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("serviceId"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.ServiceErrorResponse(d.ErrInvalidServiceId))
	}

	if err := h.serviceService.DeleteService(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.ServiceErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.ServiceDisableResponse())
}
