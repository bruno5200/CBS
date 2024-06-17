package handler

import (
	"log"

	"github.com/bruno5200/CSM/block/domain"
	p "github.com/bruno5200/CSM/block/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *blockHandler) Put(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		log.Printf("Error parsing id: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	block, err := h.BlockService.GetBlock(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	newBlock, err := domain.UnmarshalBlock(c.Body())

	if err != nil {
		log.Printf("Error unmarshalling block request: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	block.Update(newBlock)

	if err := h.BlockService.UpdateBlock(block); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	return c.Status(fiber.StatusAccepted).JSON(p.BlockSuccessResponse(block))
}
