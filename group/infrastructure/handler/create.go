package handler

import (
	"encoding/base64"
	"log"

	d "github.com/bruno5200/CSM/group/domain"
	p "github.com/bruno5200/CSM/group/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils/v2"
)

func (h *groupHandler) Post(c *fiber.Ctx) error {

	key, err := base64.StdEncoding.DecodeString(utils.CopyString(c.Get(d.API_KEY)))

	if err != nil {
		log.Printf("Error parsing key: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(d.ErrInvalidGroupServiceKey))
	}

	service, err := h.GroupService.GetServiceByKey(string(key))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	groupRequest, err := d.UnmarshalGroupRequest(c.Body())

	if err != nil {
		log.Printf("Error unmarshalling group request: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	group, err := d.NewGroup(groupRequest, service.Id)

	if err != nil {
		log.Printf("Error creating group: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	if err := h.GroupService.CreateGroup(group); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(p.GroupSuccessResponse(group))
}
