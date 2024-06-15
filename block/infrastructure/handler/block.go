package handler

import (
	d "github.com/bruno5200/CSM/block/domain"
	u "github.com/bruno5200/CSM/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *blockHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	block, err := h.BlockService.GetBlock(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	if err := u.CheckDir(d.FilesDir); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"block": block})
}
