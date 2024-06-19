package handler

import (
	"log"

	d "github.com/bruno5200/CSM/block/domain"
	p "github.com/bruno5200/CSM/block/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils/v2"
	"github.com/google/uuid"
)

func (h *blockHandler) GetByGroup(c *fiber.Ctx) error {

	groupId, err := uuid.Parse(c.Params("groupId"))

	if err != nil {
		log.Printf("Error parsing groupId: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidBlockGroupId))
	}

	blocks, err := h.BlockService.GetBlocksByGroup(groupId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.BlocksSuccessResponse(blocks))
}

func (h *blockHandler) GetByService(c *fiber.Ctx) error {

	serviceId, err := uuid.Parse(utils.CopyString(c.Get(d.HeaderServiceId)))

	if err != nil {
		log.Printf("Error parsing serviceId: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidBlockServiceId))
	}

	blocks, err := h.BlockService.GetBlocksByService(serviceId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.BlocksSuccessResponse(blocks))
}
