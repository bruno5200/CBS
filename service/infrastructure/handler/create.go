package handler

import (
	"encoding/base64"
	"log"

	d "github.com/bruno5200/CSM/service/domain"
	p "github.com/bruno5200/CSM/service/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
)

func (h *serviceHandler) Post(c *fiber.Ctx) error {

	serviceRequest, err := d.UnmarshalServiceRequest(c.Body())

	if err != nil {
		log.Printf("Error unmarshalling service request: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.ServiceErrorResponse(err))
	}

	service, err := d.NewService(serviceRequest)

	if err != nil {
		log.Printf("Error creating service: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.ServiceErrorResponse(err))
	}

	if err := h.serviceService.CreateService(service); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.ServiceErrorResponse(err))
	}

	service.Key = base64.RawStdEncoding.EncodeToString([]byte(service.Key))

	return c.Status(fiber.StatusCreated).JSON(p.ServiceSuccessResponse(service))
}
