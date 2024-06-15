package router

import (
	"fmt"

	m "github.com/bruno5200/CSM/api/middleware"
	a "github.com/bruno5200/CSM/block/application"
	"github.com/gofiber/fiber/v2"
)

func BlockRouter(app fiber.Router, h a.BlockHandler) {
	block := app.Group("api/v1/")
	block.Get("block/", h.Get)
	block.Get("block/group/:id", h.GetByGroup)
	block.Get("block/service/:id", h.GetByService)
	block.Post("upload/:groupId.:format", m.ApiKey(), func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Group: %s, Format:%s", c.Params("groupId"), c.Params("format")))
	})
	block.Post("block/:groupId.:format", m.ApiKey(), h.Json)
	block.Post("block/:groupId", m.ApiKey(), h.Post)
	block.Put("block/:id", m.ApiKey(), h.Put)
	block.Delete("block/:id", m.ApiKey(), h.Delete)
}
