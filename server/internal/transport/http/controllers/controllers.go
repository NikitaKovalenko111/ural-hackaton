package controllers

import (
	"ural-hackaton/internal/services"
	admin_controller "ural-hackaton/internal/transport/http/controllers/admin"
	event_controller "ural-hackaton/internal/transport/http/controllers/event"
	hub_controller "ural-hackaton/internal/transport/http/controllers/hub"
	mentor_controller "ural-hackaton/internal/transport/http/controllers/mentor"
	request_controller "ural-hackaton/internal/transport/http/controllers/request"
	user_controller "ural-hackaton/internal/transport/http/controllers/user"
)

type Controllers struct {
	AdminController   *admin_controller.AdminController
	UserController    *user_controller.UserController
	MentorController  *mentor_controller.MentorController
	HubController     *hub_controller.HubController
	RequestController *request_controller.RequestController
	EventController   *event_controller.EventController
}

func Init(svc *services.Services) *Controllers {
	return &Controllers{
		AdminController:   admin_controller.Init(svc.AdminService),
		UserController:    user_controller.Init(svc.UserService),
		MentorController:  mentor_controller.Init(svc.MentorService),
		HubController:     hub_controller.Init(svc.HubService),
		RequestController: request_controller.Init(svc.RequestService),
		EventController:   event_controller.Init(svc.EventService),
	}
}
