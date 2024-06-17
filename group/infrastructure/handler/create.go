package handler

import (
	"log"

	d "github.com/bruno5200/CSM/group/domain"
	p "github.com/bruno5200/CSM/group/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
)

func (h *groupHandler) Post(c *fiber.Ctx) error {

	groupRequest, err := d.UnmarshalGroupRequest(c.Body())

	if err != nil {
		log.Printf("Error unmarshalling group request: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	group, err := d.NewGroup(groupRequest)

	if err != nil {
		log.Printf("Error creating group: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	if err := h.GroupService.CreateGroup(group); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(p.GroupSuccessResponse(group))
}
