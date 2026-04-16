package http

import (
	"ural-hackaton/internal/services"
	"ural-hackaton/internal/transport/http/controllers"
)

type Http struct {
	services    *services.Services
	controllers *controllers.Controllers
}

func Init(services *services.Services) *Http {
	return &Http{
		services:    services,
		controllers: controllers.Init(services),
	}
}
