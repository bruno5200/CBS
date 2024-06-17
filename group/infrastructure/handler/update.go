package handler

import (
	"log"

	d "github.com/bruno5200/CSM/group/domain"
	"github.com/bruno5200/CSM/group/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *groupHandler) Put(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		log.Printf("Error parsing id: %s", err)
		c.Status(fiber.StatusBadRequest).JSON(presenter.GroupErrorResponse(d.ErrInvalidGroupId))
	}

	group, err := h.GroupService.GetGroup(id)

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(presenter.GroupErrorResponse(err))
	}

	newGroup, err := d.UnmarshalGroupRequest(c.Body())

	if err != nil {
		log.Printf("Error unmarshalling group request: %s", err)
		c.Status(fiber.StatusBadRequest).JSON(presenter.GroupErrorResponse(err))
	}

	group.Update(newGroup)

	if err := h.GroupService.UpdateGroup(group); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(presenter.GroupErrorResponse(err))
	}

	return c.Status(fiber.StatusAccepted).JSON(presenter.GroupSuccessResponse(group))
}
