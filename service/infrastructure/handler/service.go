package handler

import (
	"log"

	d "github.com/bruno5200/CSM/service/domain"
	p "github.com/bruno5200/CSM/service/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *serviceHandler) Get(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("serviceId"))

	if err != nil {
		log.Printf("Error parsing id: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.ServiceErrorResponse(d.ErrInvalidServiceId))
	}

	service, err := h.serviceService.GetService(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.ServiceErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.ServiceSuccessResponse(service))
}
