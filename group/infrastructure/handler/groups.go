package handler

import (
	"encoding/base64"
	"log"

	d "github.com/bruno5200/CSM/group/domain"
	p "github.com/bruno5200/CSM/group/infrastructure/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils/v2"
)

func (h *groupHandler) GetByService(c *fiber.Ctx) error {

	key, err := base64.StdEncoding.DecodeString(utils.CopyString(c.Get(d.API_KEY)))

	if err != nil {
		log.Printf("Error parsing key: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(d.ErrInvalidGroupServiceKey))
	}

	service, err := h.GroupService.GetServiceByKey(string(key))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	groups, err := h.GroupService.GetGroupsByService(service.Id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.GroupErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.GroupsSuccessResponse(groups))
}
