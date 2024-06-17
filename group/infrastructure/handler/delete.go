package handler

import (
	d "github.com/bruno5200/CSM/group/domain"
	p "github.com/bruno5200/CSM/group/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *groupHandler) Delete(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(d.ErrInvalidGroupId))
	}

	if err := h.GroupService.DeleteGroup(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.GroupDisableResponse())
}
