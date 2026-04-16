package http

import (
	"ural-hackaton/internal/docs"
	"ural-hackaton/internal/services"
	"ural-hackaton/internal/transport/http/controllers"

	"github.com/gofiber/fiber/v2"
)

type Http struct {
	services    *services.Services
	controllers *controllers.Controllers
	router      *fiber.App
}

func Init(services *services.Services, router *fiber.App) *Http {
	return &Http{
		services:    services,
		controllers: controllers.Init(services),
		router:      router,
	}
}

func (h *Http) Start() {
	h.router.Get("/swagger", func(ctx *fiber.Ctx) error {
		ctx.Type("html", "utf-8")
		return ctx.SendString(docs.SwaggerHTML)
	})

	h.router.Get("/swagger/openapi.yaml", func(ctx *fiber.Ctx) error {
		data, err := docs.OpenAPISpec.ReadFile("openapi.yaml")
		if err != nil {
			return err
		}

		ctx.Type("yaml", "utf-8")
		return ctx.Send(data)
	})

	h.controllers.AdminController.RegisterRoutes(h.router)
	h.controllers.MentorController.RegisterRoutes(h.router)
	h.controllers.UserController.RegisterRoutes(h.router)
	h.controllers.RequestController.RegisterRoutes(h.router)
	h.controllers.HubController.RegisterRoutes(h.router)
	h.controllers.EventController.RegisterRoutes(h.router)
	h.controllers.BookingController.RegisterRoutes(h.router)
}
