package http

import (
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
	h.controllers.AdminController.RegisterRoutes(h.router)
	h.controllers.MentorController.RegisterRoutes(h.router)
	h.controllers.UserController.RegisterRoutes(h.router)
	h.controllers.RequestController.RegisterRoutes(h.router)
	h.controllers.HubController.RegisterRoutes(h.router)
	h.controllers.EventController.RegisterRoutes(h.router)
}
