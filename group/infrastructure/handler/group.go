package handler

import (
	"fmt"
	"log"

	d "github.com/bruno5200/CSM/group/domain"
	p "github.com/bruno5200/CSM/group/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *groupHandler) Get(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("groupId"))

	if err != nil {
		log.Printf("Error parsing id: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(d.ErrInvalidGroupId))
	}

	groups, err := h.GroupService.GetGroup(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	for _, block := range *groups.Blocks {
		block.Url = fmt.Sprintf("%s/api/v1/block/%s", e.GetUrl(), block.Id)
	}

	return c.Status(fiber.StatusOK).JSON(p.GroupSuccessResponse(groups))
}
