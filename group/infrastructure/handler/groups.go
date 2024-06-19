package handler

import (
	"log"

	d "github.com/bruno5200/CSM/group/domain"
	p "github.com/bruno5200/CSM/group/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils/v2"
	"github.com/google/uuid"
)

func (h *groupHandler) GetByService(c *fiber.Ctx) error {

	serviceId, err := uuid.Parse(utils.CopyString(c.Get(d.HeaderServiceId)))

	if err != nil {
		log.Printf("Error parsing serviceId: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(d.ErrInvalidGroupServiceId))
	}

	groups, err := h.GroupService.GetGroupsByService(serviceId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.GroupsSuccessResponse(groups))
}
