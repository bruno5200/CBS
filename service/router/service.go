package router

import (
	a "github.com/bruno5200/CSM/service/application"
	"github.com/gofiber/fiber/v2"
)

func ServiceRouter(app fiber.Router, h a.ServiceHandler) {
	g := app.Group("api/v1/")
	g.Get("service/:serviceId", h.Get)
	g.Get("services", h.GetAll)
	g.Post("service/:serviceId", h.Post)
	g.Put("service/:id", h.Put)
	g.Delete("service/:id", h.Delete)

}
