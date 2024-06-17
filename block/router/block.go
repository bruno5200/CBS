package router

import (
	m "github.com/bruno5200/CSM/api/middleware"
	a "github.com/bruno5200/CSM/block/application"
	"github.com/gofiber/fiber/v2"
)

func BlockRouter(app fiber.Router, h a.BlockHandler) {
	b := app.Group("api/v1/")
	b.Get("block/:id.:format", h.GetParam)
	b.Get("block/:id", h.Get)
	b.Get("blocks/:groupId", h.GetByGroup)
	b.Get("blocks", m.ApiKey(), h.GetByService)
	b.Post("block/:groupId.:format", m.ApiKey(), h.PostParam)
	b.Post("block/:groupId", m.ApiKey(), h.Post)
	b.Put("block/:id", m.ApiKey(), h.Put)
	b.Delete("block/:id", m.ApiKey(), h.Delete)
}
