package router

import (
	m "github.com/bruno5200/CSM/api/middleware"
	a "github.com/bruno5200/CSM/group/application"
	"github.com/gofiber/fiber/v2"
)

func GroupRouter(app fiber.Router, h a.GroupHandler) {
	g := app.Group("api/v1/")
	g.Get("group/:groupId", m.ApiKey(), h.Get)
	g.Get("groups", m.ApiKey(), h.GetByService)
	g.Post("group", m.ApiKey(), h.Post)
	g.Put("group/:id", m.ApiKey(), h.Put)
	g.Delete("group/:id", m.ApiKey(), h.Delete)
}
