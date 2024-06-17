package handler

import (
	d "github.com/bruno5200/CSM/block/domain"
	p "github.com/bruno5200/CSM/block/infrastructure/presenter"
	u "github.com/bruno5200/CSM/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *blockHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	block, err := h.BlockService.GetBlock(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	if err := u.CheckDir(d.FilesDir); err != nil {
		return err
	}

	filePath := d.FilesDir + block.Name

	

	return c.Download(filePath, block.Name)
}

func (h *blockHandler) GetParam(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	block, err := h.BlockService.GetBlock(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	if err := u.CheckDir(d.FilesDir); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(p.BlockSuccessResponse(block))
}
