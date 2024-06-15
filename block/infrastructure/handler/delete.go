package handler

import (
	p "github.com/bruno5200/CSM/block/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *blockHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	if err := h.BlockService.DeleteBlock(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.BlockDisableResponse())
}
